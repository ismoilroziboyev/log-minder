# Log Minder

Log Minder is a logging service developed in Go (Golang) using Gin for HTTP server and MongoDB for data storage. It provides a simple HTTP request-response server for logging user activities. The service is designed to audit user actions for various applications.

## Features

- **HTTP Server**: The service provides an HTTP server to handle logging requests and responses.
- **MongoDB Integration**: User activities are stored in MongoDB for auditing purposes.
- **Auditing**: Log Minder logs user activities to track actions within applications.

## Getting Started

### Prerequisites

- Go (Golang)
- MongoDB

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/ismoilroziboyev/log-minder.git

2. Navigate to the project directory:
   ```bash
   cd log-minder

3. Install dependencies:
   ```bash
   go mod tidy

### Usage

1. Start the mongodb service.

2. Declare environment variables:
   ```bash
    MONGO_USER=admin
    MONGO_PASSWORD=admin
    MONGO_DB=database
    MONGO_ADDR=localhost:27017
    HTTP_HOST=localhost
    HTTP_PORT=35468
    ADMIN_USER=admin
    ADMIN_PASSWORD=password
   
3. Build and run log-minder application:
   ```bash
   go run main.go

The server should now be running on $HTTP_HOST:$HTTP_PORT.

### Endpoints

* POST : /v1/logs - endpoint to insert logs;
* GET: /v1/logs - endpoint to retreive logs with filter options;
* GET: /v1/swagger/index.html - swagger documentation for log-minder http server apis; 

### Future Tasks

* **Kafka Consumers**: Implement Kafka consumers to consume log messages from Kafka topics.
* **RabbitMQ Consumers**: Implement RabbitMQ consumers to consume log messages from RabbitMQ queues.

