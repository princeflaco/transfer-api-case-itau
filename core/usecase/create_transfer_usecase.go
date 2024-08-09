package usecase

import (
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

func (uc *CreateTransferUseCase) Execute(input input.TransferInput, accountId string) (*output.TransferOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, errors2.NewValidationError(err...)
	}

	if input.Amount > TransferMaxAmount {
		return nil, errors2.NewTransferMaxAmountError(input.Amount)
	}

	if accountId == input.TargetAccountId {
		return nil, errors2.NewInvalidFieldError("target_account_id", "Cannot transfer to the same account")
	}

	uc.Lock()
	defer uc.Unlock()

	accountFrom, err := uc.AccountRepo.GetById(accountId)
	if err != nil {
		return nil, err
	}

	accountTo, err := uc.AccountRepo.GetById(input.TargetAccountId)
	if err != nil {
		return nil, err
	}

	amount := util.FloatToCents(input.Amount)
	transfer := domain.NewTransfer(accountTo.Id, accountFrom.Id, amount)

	if err = accountFrom.Withdraw(amount); err != nil {
		transfer.NotSuccessful(err.Error())
		savedTransfer, err := uc.TransferRepo.Save(*transfer)
		if err != nil {
			return nil, err
		}
		transferOutput := output.NewTransferOutput(savedTransfer.Id, savedTransfer.Success, savedTransfer.Date, savedTransfer.Reason)
		return transferOutput, nil
	}

	accountTo.Deposit(amount)

	savedTransfer, err := uc.TransferRepo.Save(*transfer)
	if err != nil {
		return nil, err
	}

	_, err = uc.AccountRepo.Update(*accountFrom)

	if err != nil {
		return nil, err
	}

	_, err = uc.AccountRepo.Update(*accountTo)
	if err != nil {
		return nil, err
	}

	transferOutput := output.NewTransferOutput(savedTransfer.Id, savedTransfer.Success, savedTransfer.Date, "")

	return transferOutput, nil
}
