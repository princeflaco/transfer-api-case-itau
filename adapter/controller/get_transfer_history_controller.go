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

type GetTransferHistoryController struct {
	useCase usecase.GetTransferHistoryUseCase
}

func NewGetTransferHistoryController(useCase usecase.GetTransferHistoryUseCase) *GetTransferHistoryController {
	return &GetTransferHistoryController{useCase}
}

// Execute Get Transfer History lists an account transfer history
//
//	@Summary		List the transfer history of an account
//	@Description	Lists the transfer history in descending order of an account
//	@Tags			Transfer
//	@Produce		json
//	@Param			accountId	path		string	true	"Account ID"
//	@Success		200			{object}	[]output.TransferHistoryOutput
//	@Failure		500			{object}	dto.ErrorDTO
//	@Router			/transfers/{accountId} [get]
func (uc *GetTransferHistoryController) Execute(w http.ResponseWriter, r *http.Request) {
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
