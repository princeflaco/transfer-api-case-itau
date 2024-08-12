package service

import (
	"context"
	"go.uber.org/zap"
	"sync"
	"transfer-api/core/domain"
	"transfer-api/core/repository"
	"transfer-api/core/service/input"
	"transfer-api/core/service/output"
	"transfer-api/core/util"
)

type TransferService interface {
	Execute(ctx context.Context, input input.TransferInput) (*output.TransferOutput, error)
}

type TransferServiceImpl struct {
	TransferRepo repository.TransferRepository
	AccountRepo  repository.AccountRepository
	mu           sync.Mutex
}

func NewTransferServiceImpl(transferRepo repository.TransferRepository, accountRepo repository.AccountRepository) TransferService {
	return &TransferServiceImpl{
		TransferRepo: transferRepo,
		AccountRepo:  accountRepo,
	}
}

func (s *TransferServiceImpl) Execute(ctx context.Context, input input.TransferInput) (*output.TransferOutput, error) {
	requestId := util.GetRequestIDFromContext(ctx)
	log := util.GetLoggerFromContext(ctx).With(zap.String("request_id", requestId))

	s.mu.Lock()
	defer s.mu.Unlock()

	accountFrom, err := s.AccountRepo.GetById(input.AccountId)
	if err != nil {
		log.Error("error while retrieving account from database", zap.Error(err))
		return nil, err
	}

	log.Debug("retrieved account from database", zap.Any("account", accountFrom))

	accountTo, err := s.AccountRepo.GetById(input.TargetAccountId)
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
		savedTransfer, err := s.TransferRepo.Save(*transfer)
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

	savedTransfer, err := s.TransferRepo.Save(*transfer)
	if err != nil {
		log.Error("error while saving transfer", zap.Error(err))
		return nil, err
	}

	log.Debug("saved transfer", zap.Any("transfer", savedTransfer))

	_, err = s.AccountRepo.Update(*accountFrom)

	if err != nil {
		log.Error("error while updating account", zap.Error(err))
		return nil, err
	}

	log.Debug("updated account", zap.Any("account", accountFrom))

	_, err = s.AccountRepo.Update(*accountTo)
	if err != nil {
		log.Error("error while updating target account", zap.Error(err))
		return nil, err
	}

	log.Debug("updated target account", zap.Any("target_account", accountTo))

	transferOutput := output.NewTransferOutput(savedTransfer.Id, savedTransfer.Success, savedTransfer.Date, "")

	return transferOutput, nil
}
