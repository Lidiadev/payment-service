package domain

const (
	DepositReceivedEvent  = "payments.DepositReceived"
	DepositProcessedEvent = "payments.DepositProcessedEvent"
	DepositCompletedEvent = "payments.DepositCompletedEvent"
	DepositFailedEvent    = "payments.DepositFailedEvent"
)

type DepositReceived struct {
	//BaseEvent events.BaseEvent
	//EventID       string
	TransactionID string
	CustomerID    string
	Amount        float64
	Gateway       string
}

func (DepositReceived) Key() string { return DepositReceivedEvent }

type DepositProcessed struct {
	//BaseEvent     events.BaseEvent
	TransactionID string
	Status        TransactionStatus
	Gateway       string
}

func (DepositProcessed) Key() string { return DepositProcessedEvent }

type DepositCompleted struct {
	//BaseEvent     events.BaseEvent
	TransactionID string
	Status        TransactionStatus
	Gateway       string
}

func (DepositCompleted) Key() string { return DepositCompletedEvent }
