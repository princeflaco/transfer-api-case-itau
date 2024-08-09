package usecase

import (
	"context"
	"go.uber.org/zap"
	"sync"
	"transfer-api/core/domain"
	errors2 "transfer-api/core/errors"
	"transfer-api/core/repository"
	"transfer-api/core/usecase/input"
	"transfer-api/core/usecase/output"
	"transfer-api/core/util"
)

type CreateTransferUseCase struct {
	TransferRepo repository.TransferRepository
	AccountRepo  repository.AccountRepository
	sync.Mutex
}

const TransferMaxAmount = 10000.0

func NewCreateTransferUseCase(transferRepo repository.TransferRepository, accountRepo repository.AccountRepository) *CreateTransferUseCase {
	return &CreateTransferUseCase{
		TransferRepo: transferRepo,
		AccountRepo:  accountRepo,
	}
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

	uc.Lock()
	defer uc.Unlock()

	accountFrom, err := uc.AccountRepo.GetById(accountId)
	if err != nil {
		log.Error("error while retrieving account from database", zap.Error(err))
		return nil, err
	}

	log.Debug("retrieved account from database", zap.Any("account", accountFrom))

	accountTo, err := uc.AccountRepo.GetById(input.TargetAccountId)
	if err != nil {
		log.Error("error while retrieving target account from database", zap.Error(err))
		return nil, err
	}

	log.Debug("retrieved target account from database", zap.Any("target_account", accountTo))

	amount := util.FloatToCents(input.Amount)
	transfer := domain.NewTransfer(accountTo.Id, accountFrom.Id, amount)

	if err = accountFrom.Withdraw(amount); err != nil {
		log.Error("error while withdrawing account", zap.Error(err))
		transfer.NotSuccessful(err.Error())
		savedTransfer, err := uc.TransferRepo.Save(*transfer)
		if err != nil {
			log.Error("error while saving transfer", zap.Error(err))
			return nil, err
		}
		log.Debug("saved transfer", zap.Any("transfer", savedTransfer))
		transferOutput := output.NewTransferOutput(savedTransfer.Id, savedTransfer.Success, savedTransfer.Date, savedTransfer.Reason)
		return transferOutput, nil
	}

	accountTo.Deposit(amount)
	log.Debug("deposit amount to target account", zap.Any("amount", amount))

	savedTransfer, err := uc.TransferRepo.Save(*transfer)
	if err != nil {
		log.Error("error while saving transfer", zap.Error(err))
		return nil, err
	}

	log.Debug("saved transfer", zap.Any("transfer", savedTransfer))

	_, err = uc.AccountRepo.Update(*accountFrom)

	if err != nil {
		log.Error("error while updating account", zap.Error(err))
		return nil, err
	}

	log.Debug("updated account", zap.Any("account", accountFrom))

	_, err = uc.AccountRepo.Update(*accountTo)
	if err != nil {
		log.Error("error while updating target account", zap.Error(err))
		return nil, err
	}

	log.Debug("updated target account", zap.Any("target_account", accountTo))

	transferOutput := output.NewTransferOutput(savedTransfer.Id, savedTransfer.Success, savedTransfer.Date, "")

	return transferOutput, nil
}
