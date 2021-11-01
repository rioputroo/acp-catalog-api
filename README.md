# Project Base Alterra Academy! :star:
The ALTA Store demonstrates microservices with Go using echo server. The code implementation was inspired by port and adapter pattern or known as hexagonal:

-   **Business**<br/>Contains all the logic in domain business. Also called this as a service. All the interface of repository needed and the implementation of the service itself will be put here.
-   **Modules**<br/>Contains implementation of interfaces that defined at the business (also called as server-side adapters in hexagonal's term)
-   **Controller**<br/>Controller http handler or api (also called user-side adapters in hexagonal's term)



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


## Use of RabbitMQ in Microservices
RabbitMQ is one of the simplest freely available options for implementing messaging queues in your microservices architecture. These queue patterns can help to scale your application by communicating between various microservices.
In this project using RabbitMQ still able to communicate between various microservices


![image](https://user-images.githubusercontent.com/51318143/139615834-39f3edad-eeb4-4f19-b253-a8c8de2366c5.png)




