package controller

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"transfer-api/adapter/response"
	"transfer-api/core/usecase"
	"transfer-api/core/usecase/input"
	"transfer-api/core/util"
)

type CreateCustomerController struct {
	useCase usecase.CreateCustomerUseCase
}

func NewCreateCustomerController(useCase usecase.CreateCustomerUseCase) *CreateCustomerController {
	return &CreateCustomerController{useCase}
}

// Execute Create Customer creates a customer and account
//
//	@Summary		Create Customer
//	@Description	Creates a customer and an account
//	@Tags			Customer
//	@Accept			json
//	@Produce		json
//	@Param			customer	body		dto.CreateCustomerInput	true	"Customer and Account information"
//	@Success		201			{object}	output.CreateCustomerOutput
//	@Failure		400			{object}	dto.ErrorDTO
//	@Failure		500			{object}	dto.ErrorDTO
//	@Router			/customers [post]
func (uc *CreateCustomerController) Execute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestId := util.GetRequestIdFromHeader(r)
	ctx = context.WithValue(ctx, "request_id", requestId)

	var payload input.CreateCustomerInput
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		response.BadRequest(w, err)
		return
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			response.InternalServerError(w, err)
			return
		}
	}(r.Body)

	output, err := uc.useCase.Execute(ctx, payload)
	if err != nil {
		response.BadRequest(w, err)
		return
	}

	responseJson, err := json.Marshal(output)
	if err != nil {
		response.InternalServerError(w, err)
		return
	}

	response.Created(w, &responseJson)
}
