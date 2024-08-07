package usecase

import (
	"math"
	"transfer-api/core/domain"
	"transfer-api/core/repository"
	"transfer-api/core/usecase/input"
	"transfer-api/core/usecase/output"
)

type TransferUseCase struct {
	TransferRepo repository.TransferRepository
	AccountRepo  repository.AccountRepository
}

func NewTransferUseCase(transferRepo repository.TransferRepository, accountRepo repository.AccountRepository) *TransferUseCase {
	return &TransferUseCase{
		TransferRepo: transferRepo,
		AccountRepo:  accountRepo,
	}
}

func (uc *TransferUseCase) Execute(input input.TransferInput, accountId string) (*output.TransferOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	accountFrom, err := uc.AccountRepo.GetById(accountId)
	if err != nil {
		return nil, err
	}

	accountTo, err := uc.AccountRepo.GetById(input.TargetAccountId)
	if err != nil {
		return nil, err
	}

	amount := floatToCents(input.Amount)
	transfer := domain.NewTransfer(accountTo.Id, accountFrom.Id, amount)

	err = accountFrom.Withdraw(amount)

	if err = accountFrom.Withdraw(amount); err != nil {
		transfer.Successful(false)
		_, _ = uc.TransferRepo.SaveTransfer(*transfer)
		return nil, err
	}

	accountTo.Deposit(amount)
	transfer.Successful(true)

	savedTransfer, err := uc.TransferRepo.SaveTransfer(*transfer)
	if err != nil {
		return nil, err
	}

	_, err = uc.AccountRepo.Save(accountFrom)

	if err != nil {
		return nil, err
	}

	_, err = uc.AccountRepo.Save(accountTo)
	if err != nil {
		return nil, err
	}

	transferOutput := output.NewTransferOutput(savedTransfer.Id, savedTransfer.Success, savedTransfer.Date)

	return transferOutput, nil
}

func floatToCents(amount float64) int {
	return int(math.Round(amount * 100))
}
