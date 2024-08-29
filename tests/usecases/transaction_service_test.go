package usecases

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"payment-service/internal/application/commands"
	"payment-service/internal/application/usecases"
	"payment-service/internal/domain"
	"payment-service/internal/domain/events"
	"testing"
)

type MockEventStore struct {
	mock.Mock
}

func (m *MockEventStore) GetEvents() ([]events.Event, error) {
	return nil, nil
}

func (m *MockEventStore) SaveEvent(event events.Event) error {
	args := m.Called(event)
	return args.Error(0)
}

type MockQueue struct {
	mock.Mock
}

func (m *MockQueue) Subscribe(topic string, handler func(message interface{}) error) error {
	return nil
}

func (m *MockQueue) Publish(key string, message interface{}) error {
	args := m.Called(key, message)
	return args.Error(0)
}

func TestTransactionService_HandleDepositCommand_Success(t *testing.T) {
	// Arrange
	eventStore := new(MockEventStore)
	queue := new(MockQueue)
	service := usecases.NewTransactionService(eventStore, queue)

	cmd := commands.ProcessDepositCommand{
		CustomerID: uuid.New().String(),
		Amount:     100.0,
		GatewayID:  "gatewayA",
	}

	eventStore.On("SaveEvent", mock.AnythingOfType("events.Event")).Return(nil)
	queue.On("Publish", mock.Anything, mock.Anything).Return(nil)

	// Act
	result, err := service.HandleDepositCommand(context.Background(), cmd)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, domain.PendingStatus, result.Status)

	eventStore.AssertCalled(t, "SaveEvent", mock.AnythingOfType("events.Event"))
	queue.AssertCalled(t, "Publish", mock.AnythingOfType("string"), mock.AnythingOfType("events.Event"))
}

func TestTransactionService_HandleDepositCommand_EventStoreError(t *testing.T) {
	// Arrange
	eventStore := new(MockEventStore)
	queue := new(MockQueue)
	service := usecases.NewTransactionService(eventStore, queue)

	cmd := commands.ProcessDepositCommand{
		CustomerID: uuid.New().String(),
		Amount:     100.0,
		GatewayID:  "gatewayA",
	}

	eventStore.On("SaveEvent", mock.AnythingOfType("events.Event")).Return(errors.New("event store error"))

	// Act
	result, err := service.HandleDepositCommand(context.Background(), cmd)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to save DepositReceived event")

	queue.AssertNotCalled(t, "Publish", mock.Anything, mock.Anything)
}

func TestTransactionService_HandleDepositCommand_QueueError(t *testing.T) {
	eventStore := new(MockEventStore)
	queue := new(MockQueue)
	service := usecases.NewTransactionService(eventStore, queue)

	cmd := commands.ProcessDepositCommand{
		CustomerID: uuid.New().String(),
		Amount:     100.0,
		GatewayID:  "gatewayA",
	}

	eventStore.On("SaveEvent", mock.AnythingOfType("events.Event")).Return(nil)
	queue.On("Publish", mock.AnythingOfType("string"), mock.AnythingOfType("events.Event")).Return(errors.New("queue error"))

	// Act
	result, err := service.HandleDepositCommand(context.Background(), cmd)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to publish DepositReceived event")

	eventStore.AssertCalled(t, "SaveEvent", mock.AnythingOfType("events.Event"))
	queue.AssertCalled(t, "Publish", mock.AnythingOfType("string"), mock.AnythingOfType("events.Event"))
}
