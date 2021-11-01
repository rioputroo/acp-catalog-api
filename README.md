# PROJECT-BASE ALTERRA ACADEMY
This repository demonstrates microservices with Go using echo server. The code implementation was inspired by port and adapter pattern or known as hexagonal:

Business
Contains all the logic in domain business. Also called this as a service. All the interface of repository needed and the implementation of the service itself will be put here.
Modules
Contains implementation of interfaces that defined at the business (also called as server-side adapters in hexagonal's term)
Controller
Controller http handler or api (also called user-side adapters in hexagonal's term)

To access endpoint products :

# Use of RabbitMQ in Microservices
RabbitMQ is one of the simplest freely available options for implementing messaging queues in your microservices architecture. These queue patterns can help to scale your application by communicating between various microservices.

## Give a Star! :star:

If you like or are using this project to learn or start your solution, please give it a star. Thanks!




