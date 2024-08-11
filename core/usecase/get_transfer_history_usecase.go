package usecase

import (
	"context"
	"go.uber.org/zap"
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

func (uc *GetTransferHistoryUseCase) Execute(ctx context.Context, accountId string) ([]output.TransferHistoryOutput, error) {
	requestId := util.GetRequestIDFromContext(ctx)
	log := util.GetLoggerFromContext(ctx).With(zap.String("request_id", requestId))

	log.Info("start get transfer history use case")
	defer log.Info("end get transfer history use case")

	transfers, err := uc.TransferRepo.GetAll(accountId)
	if err != nil {
		log.Warn("error while retrieving transfer history from database", zap.Error(err))
		return []output.TransferHistoryOutput{}, err
	}

	log.Debug("retrieved transfer history from database", zap.Any("transfers", transfers))

	/*
		Ordenando slice para decrescente a partir dos timestamps.
		Utilizo a função sort.Slice() e faço uma comparação dos timestamps a partir de um parse.
	*/
	sort.Slice(transfers, func(i, j int) bool {
		dateI, errI := time.Parse(domain.TimeFormat, transfers[i].Date)
		dateJ, errJ := time.Parse(domain.TimeFormat, transfers[j].Date)

		if errI != nil || errJ != nil {
			return false
		}

		return dateI.After(dateJ)
	})
	log.Debug("sorted in descending order transfer history", zap.Any("transfers", transfers))

	history := make([]output.TransferHistoryOutput, len(transfers))
	for i, transfer := range transfers {
		amount := util.CentsToFloat64(transfer.Amount)
		history[i] = *output.NewTransferHistoryOutput(transfer.Id, transfer.AccountId, transfer.TargetAccountId, amount, transfer.Date, transfer.Success, transfer.Reason)
	}
	return history, nil
}
