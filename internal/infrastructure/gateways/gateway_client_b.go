package gateways

import (
	"bytes"
	"context"
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

var mockDepositResponseB = DepositGatewayBResponse{
	TransactionID: "mockBTx123456701",
	Status:        "success",
}

var mockWithdrawalResponseB = WithdrawalGatewayBResponse{
	TransactionID: "mockBTx9123456702",
	Status:        "success",
}

type GatewayBClientInterface interface {
	Deposit(ctx context.Context, amount float64, details map[string]string) (string, error)
	Withdraw(ctx context.Context, amount float64, details map[string]string) (string, error)
}

type GatewayBClient struct {
	endpoint   string
	timeout    time.Duration
	maxRetries int
}

func NewGatewayBClient(endpoint string, timeout time.Duration, maxRetries int) *GatewayBClient {
	return &GatewayBClient{
		endpoint: endpoint,
	}
}

func (g *GatewayBClient) Deposit(ctx context.Context, amount float64, details map[string]string) (string, error) {
	responseData, err := xml.Marshal(TransactionIDResponse{TransactionID: mockDepositResponseB.TransactionID})
	if err != nil {
		return "", err
	}

	mockHTTPResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader(responseData)),
	}

	var result struct {
		TransactionID string `xml:"transaction_id"`
	}
	if err := xml.NewDecoder(mockHTTPResponse.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.TransactionID, nil
}

func (g *GatewayBClient) Withdraw(ctx context.Context, amount float64, details map[string]string) (string, error) {
	responseData, err := xml.Marshal(TransactionIDResponse{TransactionID: mockWithdrawalResponseB.TransactionID})
	if err != nil {
		return "", err
	}

	mockHTTPResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader(responseData)),
	}

	var result struct {
		TransactionID string `xml:"transaction_id"`
	}
	if err := xml.NewDecoder(mockHTTPResponse.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.TransactionID, nil
}
