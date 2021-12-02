# microservice-template-go

go project template for microservice architecture structure.

# Description
This project is a set of ideas for implementing a microservice template in Go, I'll try to 'ilustrate' some Microservice patterns like:
- Messaging style patterns (REST)
- Reliable comunication patterns (circuit breaker)
- Security patterns 
- Cross-cutting concerns patterns

# Go Directories   

microservice-name
├── cmd <br />
|   ├── microservice-name-cli
|   |    └── main.go
|   ├── microservice-name-api
|   |    └── main.go
├── internal/
|   ├── package01
|   |    ├── file01.go
|   |    └── file02.go
|   ├── package02
|   |    ├── file01.go
|   |    └── file02.go
├── pkg/
|   ├── library01
|   |    ├── file01.go
|   |    └── file02.go
|   ├── library02
|   |    ├── file01.go
|   |    └── file02.go
├── LICENSE
└── README.md


Bibliography
https://microservices.io/patterns/
https://semver.org/
