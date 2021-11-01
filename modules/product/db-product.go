package product

import (
	"catalog/api"
	categoryController "catalog/api/category"
	productController "catalog/api/product"
	categoryService "catalog/bussiness/category"
	"catalog/bussiness/product"
	productService "catalog/bussiness/product"
	categoryRepository "catalog/modules/category"
	"encoding/json"
	"github.com/labstack/echo"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
	"log"
	"os"
)

type DbRepository struct {
	DB     *gorm.DB
	rabbit *amqp.Connection
}

type ProductTable struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"id;primaryKey;autoIncrement"`
	CategoryId  int    `json:"category_id" gorm:"category_id"`
	Name        string `json:"name" gorm:"name"`
	Price       int    `json:"price" gorm:"price"`
	Description string `json:"description" gorm:"description"`
	IsActive    bool   `json:"is_active" gorm:"is_active"`
}

//CartRabbitBody object for rabbtmq body message
type CartRabbitBody struct {
	ExchangeName      string `json:"exchange_name"`
	ExchangeType      string `json:"exchange_type"`
	PublishRoutingKey string `json:"publish_routing_key"`
	ConsumeRoutingKey string `json:"consume_routing_key"`
	QueueName         string `json:"queue_name"`
	Data              []byte `json:"data"`
}

//get field product form bussiness
func newProduct(productTemp product.Product) *ProductTable {
	return &ProductTable{
		productTemp.Model,
		productTemp.ID,
		productTemp.CategoryId,
		productTemp.Name,
		productTemp.Price,
		productTemp.Description,
		productTemp.IsActive,
	}
}

func (field *ProductTable) ToProduct() product.Product {

	var product product.Product
	product.ID = field.ID
	product.CategoryId = field.CategoryId
	product.Name = field.Name
	product.Price = field.Price
	product.Description = field.Description
	product.IsActive = field.IsActive
	return product
}

func NewProductRepository(db *gorm.DB, rabbit *amqp.Connection) *DbRepository {
	return &DbRepository{
		db,
		rabbit,
	}
}

func (temp *DbRepository) GetProductById(id int) (*product.Product, error) {

	var data ProductTable
	err := temp.DB.First(&data, id).Error
	if err != nil {
		return nil, err
	}
	product := data.ToProduct()
	return &product, nil
}

func (temp *DbRepository) GetProductsByCategoryId(categoryId int) ([]product.Product, error) {
	var data []ProductTable
	var result []product.Product

	err := temp.DB.Find(&data, ProductTable{CategoryId: categoryId}).Error

	if err != nil {
		return nil, err
	}

	for _, value := range data {
		result = append(result, value.ToProduct())
	}

	return result, nil

}

func (temp *DbRepository) GetAllProducts() ([]product.Product, error) {

	var data []ProductTable
	var result []product.Product

	err := temp.DB.Find(&data).Error
	if err != nil {
		return nil, err
	}

	for _, value := range data {
		result = append(result, value.ToProduct())
	}

	return result, nil

}

func (temp *DbRepository) CreateProduct(product product.Product) error {

	data := newProduct(product)

	err := temp.DB.Save(data).Error
	if err != nil {
		return err
	}
	return nil

}

func (temp *DbRepository) UpdateProduct(product product.Product, id int) error {

	data := newProduct(product)
	err := temp.DB.Where("id = ?", &id).Updates(&data).Error
	if err != nil {
		return err
	}
	return nil

}

func (temp *DbRepository) DeleteProduct(id int) error {

	var product ProductTable

	err := temp.DB.Delete(&product, id).Error

	if err != nil {
		return err
	}
	return nil

}

//rabbit

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (temp *DbRepository) Consume(Conn *gorm.DB) {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"rpc_queue", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	var bodyReceived map[string]interface{}
	var productTable ProductTable

	go func() {
		for d := range msgs {

			parseBody := json.Unmarshal(d.Body, &bodyReceived)

			if parseBody != nil {
				log.Printf("failed to parse body")
			}

			log.Printf("[x] Receive: %v", bodyReceived["product_id"])

			err = temp.DB.First(&productTable, bodyReceived["product_id"]).Error

			if err != nil {
				log.Printf("product not found")
			}

			productResult := productTable.ToProduct()
			dataJson, _ := json.Marshal(productResult)

			err = ch.Publish(
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "application/json",
					CorrelationId: d.CorrelationId,
					Body:          dataJson,
				})
			failOnError(err, "Failed to publish a message")

			log.Printf("[x] Sent %s", dataJson)

			d.Ack(false)
		}
	}()

	log.Printf(" [*] Awaiting RPC requests")

	prodService := productService.NewService(temp)

	prodHandler := productController.NewController(prodService)

	catRepository := categoryRepository.NewCategoryRepository(Conn)
	catService := categoryService.NewService(catRepository)
	catHandler := categoryController.NewController(catService)

	e := echo.New()
	api.HandlerApi(e, prodHandler, catHandler)

	e.Logger.Fatal(e.Start(os.Getenv("CATALOG_APP_PORT")))
	<-forever
}

func (temp *DbRepository) Publish(result product.Product) {

}

//createChannel create new channel in rabbitmq server
func (temp *DbRepository) createChannel() *amqp.Channel {
	ch, err := temp.rabbit.Channel()

	if err != nil {
		log.Printf("Failed to open a channel: %s", err)
	}

	return ch
}

//createExchange create exchange for communicate
func (temp *DbRepository) createExchange(rabbitChannel *amqp.Channel, params CartRabbitBody) {
	err := rabbitChannel.ExchangeDeclare(
		params.ExchangeName,
		params.ExchangeType,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Printf("Failed to declare an exchange: %s", err)
	}

	log.Printf("[*] Exchange created %s", params.ExchangeName)
}

//createQueue create queue for consuming messages
func (temp *DbRepository) createQueue(rabbitChannel *amqp.Channel, params CartRabbitBody) {
	_, err := rabbitChannel.QueueDeclare(
		"rpc_queue", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)

	if err != nil {
		log.Println(err, "Failed to declare a queue")
	}

	log.Printf("Usage: %s [binding_key]...", params.ConsumeRoutingKey)

	err = rabbitChannel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)

	//err = rabbitChannel.QueueBind(
	//	q.Name,                   // queue name
	//	params.ConsumeRoutingKey, // routing key
	//	params.ExchangeName,      // exchange
	//	false,
	//	nil)
	//
	//if err != nil {
	//	log.Println(err, "Failed to bind a queue")
	//}
}
