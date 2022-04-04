# gin-mongo-api

## https://dev.to/hackmamba/build-a-rest-api-with-golang-and-mongodb-gin-gonic-version-269m

# >

Representational state transfer (REST) is an architectural pattern that guides an Application programming interface(API) design and development. REST APIs have become the standard of communication between the server part of the product and its client to increase performance, scalability, simplicity, modifiability, visibility, portability, and reliability.

This post will discuss building a user management application with Golang using the Gin-gonic framework and MongoDB. At the end of this tutorial, we will learn how to structure a Gin-gonic application, build a REST API and persist our data using MongoDB.

Gin-gonic, popularly known as Gin, is an HTTP web framework written in Golang with performance and productivity support. Gin uses a custom version of HttpRouter, a lightweight, high-performance HTTP request router that navigates through API routes faster than most frameworks out there.

MongoDB is a document-based database management program used as an alternative to relational databases. MongoDB supports working with large sets of distributed data with options to store or retrieve information seamlessly.

## All Dependency

go get -u github.com/gin-gonic/gin go.mongodb.org/mongo-driver/mongo github.com/joho/godotenv github.com/go-playground/validator/v10

github.com/gin-gonic/gin is a framework for building web applications.

go.mongodb.org/mongo-driver/mongo is a driver for connecting to MongoDB.

github.com/joho/godotenv is a library for managing environment variables.

github.com/go-playground/validator/v10 is a library for validating structs and fields.

# > password hashing

https://gowebexamples.com/password-hashing/

"golang.org/x/crypto/bcrypt"

# access token 3:)

https://dev.to/joojodontoh/build-user-authentication-in-golang-with-jwt-and-mongodb-2igd

# best 3 architecture for GoLang

- Generic repository pattern (centralizes a common repository for database operations)
- Onion architecture (segregates a monolith project into controller layer, buisness layer, data access layer, domain/entity layer)
- Hexagonal architecture

## reflect.TypeOF() to get type
