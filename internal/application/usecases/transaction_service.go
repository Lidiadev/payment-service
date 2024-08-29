package usecases

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"payment-service/internal/application/commands"
	"payment-service/internal/domain"
	"payment-service/internal/domain/events"
	"payment-service/pkg"
)

type TransactionUseCase interface {
	HandleDepositCommand(ctx context.Context, cmd commands.ProcessDepositCommand) (*commands.ProcessDepositCommandResult, error)
}

type TransactionService struct {
	eventStore events.EventStore
	queue      pkg.Queue
}

func NewTransactionService(eventStore events.EventStore, queue pkg.Queue) *TransactionService {
	return &TransactionService{
		eventStore: eventStore,
		queue:      queue,
	}
}

func (s *TransactionService) HandleDepositCommand(ctx context.Context, cmd commands.ProcessDepositCommand) (*commands.ProcessDepositCommandResult, error) {
	transactionID := uuid.New().String()

	depositReceived := domain.DepositReceived{
		CustomerID:    cmd.CustomerID,
		Amount:        cmd.Amount,
		TransactionID: uuid.New().String(),
		Gateway:       cmd.GatewayID,
	}
	event := events.NewEvent(domain.DepositReceivedEvent, &depositReceived) //,cmd.Details)

	if err := s.eventStore.SaveEvent(event); err != nil {
		return nil, fmt.Errorf("failed to save DepositReceived event: %w", err)
	}

	if err := s.queue.Publish(depositReceived.Key(), event); err != nil {
		return nil, fmt.Errorf("failed to publish DepositReceived event: %w", err)
	}

	return &commands.ProcessDepositCommandResult{TransactionID: transactionID, Status: domain.PendingStatus}, nil
}
