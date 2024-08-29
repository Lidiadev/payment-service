package domain

import "strings"

type GatewayName string

const (
	GatewayAName GatewayName = "gatewayA"
	GatewayBName GatewayName = "gatewayB"
)

func (g GatewayName) ToLower() string {
	return strings.ToLower(string(g))
}
