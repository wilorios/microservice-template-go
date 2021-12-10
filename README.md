# microservice-template-go

go project template for microservice architecture structure.

# Description
This project is a set of ideas for implementing a microservice template with Go, I'll to ilustrate some Microservice patterns like:
- Messaging style patterns (REST)
- Reliable comunication patterns (circuit breaker)
- Security patterns 
- Cross-cutting concerns patterns

# Go Directories   

<pre>
microservice-name
├── cmd
|   ├── microservice-name-cli
|   |    └── main.go
|   ├── microservice-name-api
|   |    └── main.go
├── internal/
|   ├── adapters
|   |    ├── db
|   |    |    ├── file01.go
|   |    |    ├── file02.go
|   |    └── web
|   |    |    ├── file01.go
|   |    |    ├── file02.go
|   ├── application
|   |    ├── file01.go
|   |    └── file02.go
|   ├── configurations
|   |    ├── file01.go
|   |    └── file02.go
|   ├── model
|   |    ├── file01.go
|   |    └── file02.go
├── pkg/
|   ├── library01
|   |    ├── file01.go
|   |    └── file02.go
|   ├── library02
|   |    ├── file01.go
|   |    └── file02.go
├── go.mod
├── go.sum
└── README.md
</pre>

Bibliography
- https://microservices.io/patterns/
- https://semver.org/
