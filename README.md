## Payment Gateway Microservice

### High-Level Architecture Overview
This payment gateway microservice is designed to integrate with multiple payment gateways and manage deposits. The architecture follows clean architecture principles.

#### Key components:

- API Layer: Handles HTTP requests using the Gin framework.
- Application Layer: Contains use cases and commands for business logic.
- Domain Layer: Defines core business logic and entities.
- Infrastructure Layer: Implements external interfaces (gateways, databases, etc.).
- Gateway Factory: Allows easy integration of new payment gateways.
- Resilience Patterns: Implements circuit breaker and retry mechanisms.

## Event-Driven Architecture
The microservice employs an event-driven architecture to enhance scalability, decouple components and improve fault tolerance. 

### How It Works

1. **Event Generation**: When an action occurs (e.g. a deposit request), the system generates an event (e.g. DepositReceived).
2. **Event Publishing**: The event is published to a queue (currently an in-memory queue, but it can be replaced with a distributed message broker like Kafka or RabbitMQ for production use).
3. **Event Processing**: A worker (implemented in the workers package) subscribes to the queue and processes events asynchronously.

## Event Processing
The event processing flow in this microservice works as follows:

1. **Event Worker**: The EventWorker in the workers package is responsible for processing events from the queue.
2. **Event Handling**: For each event, the worker:
- Determines the event type (e.g. DepositReceived)
- Retrieves the appropriate payment gateway using the GatewayFactory
- Calls the relevant method on the gateway (e.g., Deposit)
- Based on the result, generates and publishes a new event (e.g., DepositProcessed or DepositFailed)


3. **Concurrent Processing**: The worker can process multiple events concurrently, improving throughput.

### Failure Handling
The system is designed to be resilient to failures at various levels:
1. **Event Processing Failures**:
- If an event fails to process, it's moved to a Dead Letter Queue (DLQ).
- The system logs the failure for monitoring and debugging.
- Events in the DLQ can be retried later or manually processed.

2. **Gateway Failures**:
- The system uses the Hystrix library for implementing the Circuit Breaker pattern.
- If a gateway consistently fails, the circuit opens, preventing further calls to the failing gateway.
- During this time, new requests are immediately rejected or routed to a fallback method.
- After a cooldown period, the circuit half-opens to test if the gateway has recovered.
- For transient failures, the system can retry operations with exponential backoff.


3. System Crashes:
- Events are persisted before processing, ensuring they're not lost if the system crashes.
- Upon restart, unprocessed events can be reloaded and processed.

#### Saving Events:
When an event is generated, it's saved to the event store using the SaveEvent method.
Each event includes metadata such as event type, timestamp, and associated data.
For production use, we should ensure events are saved durably before acknowledging their receipt.

## Resilience to Failures
The microservice is designed to be resilient to various types of failures:

- Circuit Breaker Pattern
- Retry Mechanism
- Dead Letter Queue (DLQ):
Messages that repeatedly fail processing are moved to a DLQ. This prevents problematic messages from blocking the processing of other messages.
- Asynchronous Processing:
By processing events asynchronously, the system can continue to accept new requests even if some are failing. This improves overall system availability and responsiveness.

Monitoring and Alerts:

The failure is logged and monitoring systems should be configured.


By implementing these resilience patterns and failure handling mechanisms, the microservice can maintain a high level of availability and reliability, even in the face of various failure scenarios.

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
- Implement Withdrawal and Callback
- Process all types of events
- Add more robust error handling
- Process failed events
- Integrate with a DB and message broker
- Decouple from `time.Now()` using the `utils.Providers` interface
- Add more UTs and integration tests.
