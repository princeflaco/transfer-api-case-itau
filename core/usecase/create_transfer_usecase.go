package usecase

import (
	"context"
	"go.uber.org/zap"
	"sync"
	errors2 "transfer-api/core/errors"
	"transfer-api/core/service"
	"transfer-api/core/service/dto"
	"transfer-api/core/usecase/input"
	"transfer-api/core/usecase/output"
	"transfer-api/core/util"
)

type CreateTransferUseCase struct {
	TransferMaxAmount   float64
	TransferWorkerCount int
	Service             service.TransferService
	Queue               chan dto.TransferRequest
	wg                  sync.WaitGroup
}

type TransferConfig struct {
	MaxAmount   float64
	WorkerCount int
}

func NewCreateTransferUseCase(service service.TransferService, config TransferConfig) *CreateTransferUseCase {
	useCase := &CreateTransferUseCase{
		TransferMaxAmount:   config.MaxAmount,
		TransferWorkerCount: config.WorkerCount,
		Service:             service,
		Queue:               make(chan dto.TransferRequest, 100),
	}
	useCase.startWorkers()
	return useCase
}

func (uc *CreateTransferUseCase) startWorkers() {
	for i := 0; i < uc.TransferWorkerCount; i++ {
		uc.wg.Add(1)
		go uc.transferWorker()
	}
}

func (uc *CreateTransferUseCase) transferWorker() {
	defer uc.wg.Done()

	for req := range uc.Queue {
		req.Result <- uc.Service.Execute(req)
		close(req.Result)
	}
}

func (uc *CreateTransferUseCase) Execute(ctx context.Context, input input.CreateTransferInput, accountId string) (*output.CreateTransferOutput, error) {
	requestId := util.GetRequestIDFromContext(ctx)
	log := util.GetLoggerFromContext(ctx).With(zap.String("request_id", requestId))

	log.Info("start create transfer use case")
	defer log.Info("end create transfer use case")

	if err := input.Validate(); err != nil {
		err := errors2.NewValidationError(err...)
		log.Error("error while validating dto", zap.Error(err))
		return nil, err
	}

	if input.Amount > uc.TransferMaxAmount {
		err := errors2.NewTransferMaxAmountError(input.Amount)
		log.Error("amount is too high", zap.Error(err))
		return nil, err
	}

	if accountId == input.TargetAccountId {
		err := errors2.NewInvalidFieldError("target_account_id", "Cannot transfer to the same account")
		log.Error("impossible to transfer to own account", zap.Error(err))
		return nil, err
	}

	resultChan := make(chan dto.TransferResult)
	uc.EnqueueTransfer(ctx, input, resultChan)

	result := <-resultChan

	if result.Error != nil {
		return nil, result.Error
	}

	return result.Output, nil
}

func (uc *CreateTransferUseCase) EnqueueTransfer(ctx context.Context, input input.CreateTransferInput, resultChan chan dto.TransferResult) {
	uc.Queue <- *dto.NewTransferRequest(ctx, input, resultChan)
}

func (uc *CreateTransferUseCase) Shutdown() {
	close(uc.Queue)
	uc.wg.Wait()
}
