package product

import (
	"catalog/bussiness/product"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
	"log"
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

func (temp *DbRepository) GetAllProducts(categoryId int) ([]product.Product, error) {

	var data []ProductTable
	var result []product.Product

	if &categoryId != nil {
		err := temp.DB.Find(&data, ProductTable{CategoryId: categoryId}).Error
		if err != nil {
			return nil, err
		}
	} else {
		err := temp.DB.Find(&data).Error
		if err != nil {
			return nil, err
		}
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
func (temp *DbRepository) Consume() {
	var rabbitBody CartRabbitBody

	rabbitBody.ExchangeName = "add_to_cart"
	rabbitBody.ExchangeType = "topic"
	rabbitBody.PublishRoutingKey = "product.product.read"
	rabbitBody.ConsumeRoutingKey = "cart.item.added"
	rabbitBody.QueueName = "cart_product_queue"

	rabbitChannel := temp.createChannel()

	temp.createExchange(rabbitChannel, rabbitBody)
	temp.createQueue(rabbitChannel, rabbitBody)

	msgs, err := rabbitChannel.Consume(
		rabbitBody.QueueName, // queue
		"",                   // consumer
		true,                 // auto ack
		false,                // exclusive
		false,                // no local
		false,                // no wait
		nil,                  // args
	)

	if err != nil {
		log.Println(err, "Failed to register a consumer")
	}

	//forever := make(chan bool)

	var productId float64
	var productTable ProductTable

	//go func() {
	var bodyReceived map[string]interface{}

	for d := range msgs {
		parseBody := json.Unmarshal(d.Body, &bodyReceived)

		if parseBody != nil {
			log.Printf("failed to parse body")
		}

		log.Printf("[x] Receive: %v", bodyReceived["product_id"])
		productId = bodyReceived["product_id"].(float64)

		log.Println(productId)

		err = temp.DB.First(&productTable, productId).Error

		if err != nil {
			log.Printf("product not found")
		}

		productResult := productTable.ToProduct()
		dataJson, _ := json.Marshal(productResult)

		err = rabbitChannel.Publish(
			rabbitBody.ExchangeName,
			rabbitBody.PublishRoutingKey,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        dataJson},
		)

		if err != nil {
			log.Printf("Failed to publish message: %s", err)
		}

		log.Printf("[x] Sent %s", string(dataJson))
	}

	//}()

	log.Printf("[*] Waiting for any request. To exit press CTRL+C")

	//defer rabbitChannel.Close()

	//<-forever
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
	q, err := rabbitChannel.QueueDeclare(
		params.QueueName, // name
		false,            // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)

	if err != nil {
		log.Println(err, "Failed to declare a queue")
	}

	log.Printf("Usage: %s [binding_key]...", params.ConsumeRoutingKey)

	err = rabbitChannel.QueueBind(
		q.Name,                   // queue name
		params.ConsumeRoutingKey, // routing key
		params.ExchangeName,      // exchange
		false,
		nil)

	if err != nil {
		log.Println(err, "Failed to bind a queue")
	}
}
