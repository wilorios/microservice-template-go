# microservice-template-go

go project template for microservice architecture structure.

# Description
This project is a set of ideas for implementing a microservice template in Go, I'll try to 'ilustrate' some Microservice patterns like:
- Messaging style patterns (REST)
- Reliable comunication patterns (circuit breaker)
- Security patterns 
- Cross-cutting concerns patterns

# Go Directories   

microservice-name<br/>
├── cmd<br/>
|   ├── microservice-name-cli<br/>
|   |    └── main.go<br/>
|   ├── microservice-name-api<br/>
|   |    └── main.go<br/>
├── internal/<br/>
|   ├── package01<br/>
|   |    ├── file01.go<br/>
|   |    └── file02.go<br/>
|   ├── package02<br/>
|   |    ├── file01.go<br/>
|   |    └── file02.go<br/>
├── pkg/<br/>
|   ├── library01<br/>
|   |    ├── file01.go<br/>
|   |    └── file02.go<br/>
|   ├── library02<br/>
|   |    ├── file01.go<br/>
|   |    └── file02.go<br/>
├── LICENSE<br/>
└── README.md<br/>


Bibliography
https://microservices.io/patterns/
https://semver.org/
