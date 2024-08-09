package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"transfer-api/adapter/response"
	"transfer-api/core/usecase"
)

type GetCustomerController struct {
	useCase usecase.GetCustomerUseCase
}

func NewGetCustomerController(useCase usecase.GetCustomerUseCase) *GetCustomerController {
	return &GetCustomerController{useCase}
}

func (uc *GetCustomerController) Execute(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			response.BadRequest(w, err)
		}
	}(r.Body)

	accountId := r.URL.Query().Get("accountId")

	output, err := uc.useCase.Execute(accountId)
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
