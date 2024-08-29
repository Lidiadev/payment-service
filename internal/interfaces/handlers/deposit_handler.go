package handlers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"payment-service/internal/application/commands"
	"payment-service/internal/application/usecases"
	"payment-service/internal/infrastructure/logging"
	"payment-service/internal/interfaces/models"
	"payment-service/internal/interfaces/utils"
	"payment-service/internal/interfaces/validators"
)

type DepositHandler struct {
	transactionUseCase usecases.TransactionUseCase
	validator          validators.ValidatorInterface
	logger             logging.LoggerInterface
}

func NewDepositHandler(transactionUseCase usecases.TransactionUseCase, validator validators.ValidatorInterface, logger logging.LoggerInterface) *DepositHandler {
	return &DepositHandler{
		transactionUseCase: transactionUseCase,
		validator:          validator,
		logger:             logger,
	}
}

// @Summary Deposit money into an account
// @Description This will deposit a specific amount into a user's account using a payment gateway.
// @Tags Deposits
// @Accept  json
// @Produce  json
// @Param request body models.DepositRequest true "Deposit Details"
// @Success 200 {object} models.DepositResponse
// @Failure 400 {string} string "Invalid request body"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/deposit [post]
func (h *DepositHandler) HandleDeposit(c *gin.Context) {
	var request models.DepositRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		h.logger.Error(err.Error())
		utils.SendErrorResponse(c, http.StatusBadRequest, utils.ErrInvalidRequestBody)
		return
	}

	if err := h.validator.Validate(request); err != nil {
		validationErr := err.(validator.ValidationErrors)
		utils.SendErrorResponse(c, http.StatusBadRequest, fmt.Sprintf(utils.ErrValidationRequest, validationErr))
		return
	}

	req := commands.ProcessDepositCommand{
		GatewayID:  request.GatewayID,
		CustomerID: request.CustomerID,
		Amount:     request.Amount,
		Details:    request.Details,
	}

	commandResponse, err := h.transactionUseCase.HandleDepositCommand(context.Background(), req)
	if err != nil {
		h.logger.Error(err.Error())
		utils.SendErrorResponse(c, http.StatusInternalServerError, utils.ErrDepositFailed)
		return
	}

	response := models.DepositResponse{
		TransactionID: commandResponse.TransactionID,
		Status:        commandResponse.Status.String(),
	}

	//h.logger.LogTransaction(enum.DepositTransactionType, request, response)

	c.JSON(http.StatusOK, response)
}
