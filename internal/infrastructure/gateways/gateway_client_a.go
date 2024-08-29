package gateways

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"payment-service/internal/infrastructure/resilience"
)

var mockDepositResponseA = DepositGatewayAResponse{
	TransactionID: "mockATx100001",
	Status:        "success",
}

var mockWithdrawalResponseA = WithdrawalGatewayAResponse{
	TransactionID: "mockATx100002",
	Status:        "success",
}

type GatewayAClientInterface interface {
	Deposit(ctx context.Context, amount float64, details map[string]string) (string, error)
	Withdraw(ctx context.Context, amount float64, details map[string]string) (string, error)
}

type GatewayAClient struct {
	endpoint   string
	hystrixCmd string
}

func NewGatewayAClient(endpoint string, hystrixCmd string) *GatewayAClient {
	return &GatewayAClient{
		endpoint:   endpoint,
		hystrixCmd: hystrixCmd,
	}
}

func (g *GatewayAClient) Deposit(ctx context.Context, amount float64, details map[string]string) (string, error) {
	reqBody, err := json.Marshal(map[string]interface{}{
		"amount":  amount,
		"details": details,
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", g.endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	fallbackFunc := func(err error) error {
		return fmt.Errorf("fallback triggered for deposit: %w", err)
	}

	resp, err := g.doRequestWithHystrix(ctx, req, fallbackFunc)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var depositResponse DepositGatewayAResponse
	err = json.Unmarshal(body, &depositResponse)
	if err != nil {
		return "", err
	}

	if depositResponse.Status != "success" {
		return "", errors.New("deposit failed")
	}

	return depositResponse.TransactionID, nil
}

func (g *GatewayAClient) Withdraw(ctx context.Context, amount float64, details map[string]string) (string, error) {
	responseData, err := json.Marshal(mockWithdrawalResponseA)
	if err != nil {
		return "", err
	}

	mockHTTPResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader(responseData)),
	}

	var withdrawalResponse WithdrawalGatewayAResponse
	if err := json.NewDecoder(mockHTTPResponse.Body).Decode(&withdrawalResponse); err != nil {
		return "", err
	}

	return withdrawalResponse.TransactionID, nil
}

func (g *GatewayAClient) doRequestWithHystrix(ctx context.Context, req *http.Request, fallbackFunc func(error) error) (*http.Response, error) {
	var resp *http.Response

	runFunc := func() error {

		// Mock data instead of calling the http client
		responseData, err := json.Marshal(mockDepositResponseA)
		if err != nil {
			return err
		}

		mockHTTPResponse := &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewReader(responseData)),
		}

		resp = mockHTTPResponse
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			return fmt.Errorf("non-successful status code: %d", resp.StatusCode)
		}
		return nil
	}

	err := resilience.RunWithHystrix(g.hystrixCmd, runFunc, fallbackFunc)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
