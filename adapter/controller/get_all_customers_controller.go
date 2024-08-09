package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"transfer-api/adapter/response"
	"transfer-api/core/usecase"
)

type GetAllCustomersController struct {
	useCase usecase.GetAllCustomersUseCase
}

func NewGetAllCustomersController(useCase usecase.GetAllCustomersUseCase) *GetAllCustomersController {
	return &GetAllCustomersController{useCase}
}

func (uc *GetAllCustomersController) Execute(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			response.BadRequest(w, err)
		}
	}(r.Body)

	output, err := uc.useCase.Execute()
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
