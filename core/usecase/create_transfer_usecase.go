package usecase

import (
	"context"
	"go.uber.org/zap"
	"sync"
	errors2 "transfer-api/core/errors"
	"transfer-api/core/service"
	"transfer-api/core/service/input"
	"transfer-api/core/service/output"
	"transfer-api/core/util"
)

type (
	CreateTransferUseCase struct {
		Service service.TransferService
		Queue   chan TransferRequest
		wg      sync.WaitGroup
	}
	TransferRequest struct {
		Context context.Context
		Input   input.TransferInput
		Result  chan TransferResult
	}
	TransferResult struct {
		Output *output.TransferOutput
		Error  error
	}
)

const (
	TransferMaxAmount = 10000.0
	WorkerCount       = 10
)

func NewCreateTransferUseCase(service service.TransferService) *CreateTransferUseCase {
	useCase := &CreateTransferUseCase{
		Service: service,
		Queue:   make(chan TransferRequest, 100),
	}
	useCase.startWorkers()
	return useCase
}

func (uc *CreateTransferUseCase) startWorkers() {
	for i := 0; i < WorkerCount; i++ {
		uc.wg.Add(1)
		go uc.transferWorker()
	}
}

func (uc *CreateTransferUseCase) transferWorker() {
	defer uc.wg.Done()

	for req := range uc.Queue {
		o, err := uc.Service.Execute(req.Context, req.Input)
		req.Result <- TransferResult{
			Output: o,
			Error:  err,
		}
		close(req.Result)
	}
}

func (uc *CreateTransferUseCase) EnqueueTransfer(ctx context.Context, input input.TransferInput, resultChan chan TransferResult) {
	uc.Queue <- TransferRequest{Input: input, Context: ctx, Result: resultChan}
}

func (uc *CreateTransferUseCase) Shutdown() {
	close(uc.Queue)
	uc.wg.Wait()
}

func (uc *CreateTransferUseCase) Execute(ctx context.Context, input input.TransferInput, accountId string) (*output.TransferOutput, error) {
	requestId := util.GetRequestIDFromContext(ctx)
	log := util.GetLoggerFromContext(ctx).With(zap.String("request_id", requestId))

	log.Info("start create transfer use case")
	defer log.Info("end create transfer use case")

	if err := input.Validate(); err != nil {
		err := errors2.NewValidationError(err...)
		log.Error("error while validating input", zap.Error(err))
		return nil, err
	}

	if input.Amount > TransferMaxAmount {
		err := errors2.NewTransferMaxAmountError(input.Amount)
		log.Error("amount is too high", zap.Error(err))
		return nil, err
	}

	if accountId == input.TargetAccountId {
		err := errors2.NewInvalidFieldError("target_account_id", "Cannot transfer to the same account")
		log.Error("impossible to transfer to own account", zap.Error(err))
		return nil, err
	}

	resultChan := make(chan TransferResult)
	uc.EnqueueTransfer(ctx, input, resultChan)

	result := <-resultChan

	if result.Error != nil {
		return nil, result.Error
	}

	return result.Output, nil
}
