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

type CreateTransferController struct {
	useCase *usecase.CreateTransferUseCase
}

func NewCreateTransferController(useCase *usecase.CreateTransferUseCase) *CreateTransferController {
	return &CreateTransferController{useCase}
}

// Execute Create Transfer creates a transfer between accounts
//
//	@Summary		Transfer amount
//	@Description	Transfer an amount between accounts
//	@Tags			Transfer
//	@Accept			json
//	@Produce		json
//	@Param			accountId	path		string				true	"Account ID"
//	@Param			transfer	body		dto.CreateTransferInput	true	"Transfer information"
//	@Success		201			{object}	output.TransferOutput
//	@Failure		400			{object}	dto.ErrorDTO
//	@Failure		404			{object}	dto.ErrorDTO
//	@Failure		500			{object}	dto.ErrorDTO
//	@Router			/transfers/{accountId} [post]
func (uc *CreateTransferController) Execute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestId := util.GetRequestIdFromHeader(r)
	ctx = context.WithValue(ctx, "request_id", requestId)

	var payload input.CreateTransferInput
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

	output, err := uc.useCase.Execute(ctx, payload, accountId)
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
