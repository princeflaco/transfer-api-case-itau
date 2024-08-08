package usecase

import (
	"sort"
	"time"
	"transfer-api/core/domain"
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

	sort.Slice(transfers, func(i, j int) bool {
		dateI, errI := time.Parse(domain.TimeFormat, transfers[i].Date)
		dateJ, errJ := time.Parse(domain.TimeFormat, transfers[j].Date)

		if errI != nil || errJ != nil {
			return false
		}

		return dateI.After(dateJ)
	})

	history := make([]output.TransferHistoryOutput, len(transfers))
	for i, transfer := range transfers {
		amount := util.CentsToFloat64(transfer.Amount)
		history[i] = *output.NewTransferHistoryOutput(transfer.Id, transfer.TargetAccountId, amount, transfer.Date, transfer.Success)
	}
	return history, nil
}
