package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"transfer-api/adapter/response"
	"transfer-api/core/usecase"
	"transfer-api/core/usecase/input"
)

type CreateTransferController struct {
	useCase *usecase.CreateTransferUseCase
}

func NewCreateTransferController(useCase *usecase.CreateTransferUseCase) *CreateTransferController {
	return &CreateTransferController{useCase}
}

func (uc *CreateTransferController) Execute(w http.ResponseWriter, r *http.Request) {
	var payload input.TransferInput
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		response.BadRequest(w, err)
		return
	}

	accountId := r.URL.Query().Get("accountId")

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			response.InternalServerError(w, err)
			return
		}
	}(r.Body)

	output, err := uc.useCase.Execute(payload, accountId)
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
