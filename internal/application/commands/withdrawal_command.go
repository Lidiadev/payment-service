package commands

type ProcessWithdrawalCommand struct {
	Amount  float64
	Details map[string]string
	Gateway string
}
