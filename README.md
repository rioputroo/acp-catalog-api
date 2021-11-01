# acp-catalog-api
This is Catalog API for ACP Final Project using Golang

GOLANG REST API build using echo server.

The code implementation was inspired by port and adapter pattern or known as hexagonal:

Business
Contains all the logic in domain business. Also called this as a service. All the interface of repository needed and the implementation of the service itself will be put here.
Modules
Contains implementation of interfaces that defined at the business (also called as server-side adapters in hexagonal's term)
Controller
Controller http handler or api (also called user-side adapters in hexagonal's term)

To access endpoint products :


