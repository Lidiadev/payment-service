package domain

type TransactionStatus string

const (
	UnknownStatus   TransactionStatus = ""
	PendingStatus   TransactionStatus = "pending"
	ProcessedStatus TransactionStatus = "processed"
	CompletedStatus TransactionStatus = "completed"
	FailedStatus    TransactionStatus = "failed"
)

func (s TransactionStatus) String() string {
	switch s {
	case PendingStatus, ProcessedStatus, CompletedStatus, FailedStatus:
		return string(s)
	default:
		return ""
	}
}

func ToTransactionStatus(status string) TransactionStatus {
	switch status {
	case PendingStatus.String():
		return PendingStatus
	case ProcessedStatus.String():
		return ProcessedStatus
	case CompletedStatus.String():
		return CompletedStatus
	case FailedStatus.String():
		return FailedStatus
	default:
		return UnknownStatus
	}
}
