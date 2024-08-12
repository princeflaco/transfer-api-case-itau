package dto

import (
	"context"
	"transfer-api/core/usecase/input"
	"transfer-api/core/usecase/output"
)

type TransferRequest struct {
	Context context.Context
	Input   input.CreateTransferInput
	Result  chan TransferResult
}

func NewTransferRequest(ctx context.Context, input input.CreateTransferInput, resultChan chan TransferResult) *TransferRequest {
	return &TransferRequest{
		Context: ctx,
		Input:   input,
		Result:  resultChan,
	}
}

type TransferResult struct {
	Output *output.CreateTransferOutput
	Error  error
}

func NewTransferResult(output *output.CreateTransferOutput, err error) TransferResult {
	return TransferResult{
		Output: output,
		Error:  err,
	}
}
