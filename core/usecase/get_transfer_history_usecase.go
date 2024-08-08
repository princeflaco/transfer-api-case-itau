package usecase

import (
	"transfer-api/core/repository"
	"transfer-api/core/usecase/output"
	"transfer-api/core/util"
)

type GetTransferHistoryUseCase struct {
	TransferRepo repository.TransferRepository
}

func NewGetTransferHistoryUseCase(transferRepository repository.TransferRepository) *GetTransferHistoryUseCase {
	return &GetTransferHistoryUseCase{
		TransferRepo: transferRepository,
	}
}

func (uc *GetTransferHistoryUseCase) Execute(accountId string) ([]output.TransferHistoryOutput, error) {
	transfers, err := uc.TransferRepo.GetAll(accountId)
	if err != nil {
		return []output.TransferHistoryOutput{}, err
	}
	var outputs []output.TransferHistoryOutput
	for _, transfer := range transfers {
		amount := util.CentsToFloat64(transfer.Amount)
		transferOutput := output.NewTransferHistoryOutput(transfer.Id, transfer.TargetAccountId, amount, transfer.Date, transfer.Success)
		outputs = append(outputs, *transferOutput)
	}
	return outputs, nil
}
