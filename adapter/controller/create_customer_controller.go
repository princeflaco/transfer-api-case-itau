package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"transfer-api/adapter/response"
	"transfer-api/core/usecase"
	"transfer-api/core/usecase/input"
)

type CreateCustomerController struct {
	useCase usecase.CreateCustomerUseCase
}

func NewCreateCustomerController(useCase usecase.CreateCustomerUseCase) *CreateCustomerController {
	return &CreateCustomerController{useCase}
}

func (uc *CreateCustomerController) Execute(w http.ResponseWriter, r *http.Request) {
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

	output, err := uc.useCase.Execute(payload)
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
