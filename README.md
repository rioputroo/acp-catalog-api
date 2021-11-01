# Project Base Alterra Academy! :star:
The ALTA Store demonstrates microservices with Go using echo server. The code implementation was inspired by port and adapter pattern or known as hexagonal:

The hexagonal architecture is based on three principles and techniques:

- Explicitly separate User-Side, Business Logic, and Server-Side
- Dependencies are going from User-Side and Server-Side to the Business Logic
- We isolate the boundaries by using Ports and Adapters

![image](https://user-images.githubusercontent.com/51318143/139618165-bdaeb6d7-dbf5-4b6c-bf27-3508be3f1dc7.png)

<br>
<br>
## Use of RabbitMQ in Microservices
RabbitMQ is one of the simplest freely available options for implementing messaging queues in your microservices architecture. These queue patterns can help to scale your application communicating between microservices.
In this project using RabbitMQ to able to communicate between various microservices



![image](https://user-images.githubusercontent.com/51318143/139615834-39f3edad-eeb4-4f19-b253-a8c8de2366c5.png)




## Data initialization

To describe about how port and adapter interaction (separation concerned), this example will have two databases supported. There are MySQL using gorm as library.




## How To Consume The API

	//list about iam auth (auth management)
	POST Method "/login", route to login user
	POST Method "/register", route to register user
  
	//list about catalog product API
	GET Method "/catalog/products", to get all product
	GET Method "/catalog/product/:productId", to get product by id
	GET Method "/catalog/filterproduct/?categoryId=", to get products by category id with query params
	POST Method "/catalog/product", to create new product
	PUT Method "/catalog/product/:productId", to update product by id 
	DELETE Method "/catalog/product/:productId", to delete product
  
	//list about order PAI
	GET Method "/order/cart", to get cart
	POST Method "/order/cart", to add item to cart
	GET Method "/order/checkout", to send checkout







