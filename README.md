## Payment Gateway Microservice

### High-Level Architecture Overview
This payment gateway microservice is designed to integrate with multiple payment gateways and manage deposits. The architecture follows clean architecture principles.

#### Key components:

- API Layer: Handles HTTP requests using the Gin framework.
- Application Layer: Contains use cases and commands for business logic.
- Domain Layer: Defines core business logic and entities.
- Infrastructure Layer: Implements external interfaces (gateways, databases, etc.).
- Event-Driven Architecture: Uses events and queues for asynchronous processing.
- Gateway Factory: Allows easy integration of new payment gateways.
- Resilience Patterns: Implements circuit breaker and retry mechanisms.

## Building and Running the Service
### Prerequisites
- Docker Compose 

### Building the Service

1. Clone the repository:
```
   git clone https://github.com/Lidiadev/payment-service.git
   cd payment-service
```

2. Running the Service

Run using Docker:

```
docker compose up --build
```

3. Access the Swagger Documentation

The API documentation is generated using Swaggo and is available at:

```
http://localhost:8080/swagger/index.html
```


### Testing the Service

Run unit tests:

```
go test ./...
```


Note: The current implementation uses an in-memory queue and an in-memory db.

### API Documentation

Endpoints include:

- POST /v1/deposit: Initiate a deposit request.

Example Request for Deposit
```
curl -X POST "http://localhost:8080/v1/deposit" \
-H "Content-Type: application/json" \
-d '{
  "amount": 100.0,
  "gateway": "gatewayA",
  "customer_id": "12345"
}'
```


### Adding New Payment Gateways
To add a new payment gateway:

1. Create a new client in the gateways package.
2. Implement the PaymentGatewayStrategy interface for the new gateway.
3. Register the new gateway in the GatewayFactory in main.go.

## Future Improvements
Add more robust error handling and retries for failed events.
