package controller

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"transfer-api/adapter/response"
	"transfer-api/core/usecase"
	"transfer-api/core/util"
)

type GetCustomerController struct {
	useCase usecase.GetCustomerUseCase
}

func NewGetCustomerController(useCase usecase.GetCustomerUseCase) *GetCustomerController {
	return &GetCustomerController{useCase}
}

// Execute Get Customer retrieves a customer
//
//	@Summary		Get Customer By Account
//	@Description	Retrieves a Customer by his account_id
//	@Tags			Customer
//	@Produce		json
//	@Param			accountId	path		string	true	"Account ID"
//	@Success		200			{object}	output.GetCustomerOutput
//	@Failure		404			{object}	dto.ErrorDTO
//	@Failure		500			{object}	dto.ErrorDTO
//	@Router			/customers/{accountId} [get]
func (uc *GetCustomerController) Execute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestId := util.GetRequestIdFromHeader(r)
	ctx = context.WithValue(ctx, "request_id", requestId)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			response.BadRequest(w, err)
		}
	}(r.Body)

	accountId := r.URL.Query().Get("accountId")

	output, err := uc.useCase.Execute(ctx, accountId)
	if err != nil {
		response.BadRequest(w, err)
		return
	}

	responseJson, err := json.Marshal(output)
	if err != nil {
		response.InternalServerError(w, err)
		return
	}

	response.Ok(w, responseJson)
}
