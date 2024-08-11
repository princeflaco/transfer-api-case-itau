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

type GetAllCustomersController struct {
	useCase usecase.GetAllCustomersUseCase
}

func NewGetAllCustomersController(useCase usecase.GetAllCustomersUseCase) *GetAllCustomersController {
	return &GetAllCustomersController{useCase}
}

// Execute Get All Customers lists all customers
//
//	@Summary		List Customers
//	@Description	List all existent customers
//	@Tags			Customer
//	@Produce		json
//	@Success		200	{object}	[]output.GetCustomerOutput
//	@Failure		500	{object}	dto.ErrorDTO
//	@Router			/customers [get]
func (uc *GetAllCustomersController) Execute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestId := util.GetRequestIdFromHeader(r)
	ctx = context.WithValue(ctx, "request_id", requestId)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			response.BadRequest(w, err)
		}
	}(r.Body)

	output, err := uc.useCase.Execute(ctx)
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
