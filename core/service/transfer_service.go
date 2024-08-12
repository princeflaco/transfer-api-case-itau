package service

import (
	"go.uber.org/zap"
	"sync"
	"transfer-api/core/domain"
	"transfer-api/core/repository"
	"transfer-api/core/service/dto"
	"transfer-api/core/usecase/output"
	"transfer-api/core/util"
)

type TransferService interface {
	Execute(request dto.TransferRequest) dto.TransferResult
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

func (s *TransferServiceImpl) Execute(request dto.TransferRequest) dto.TransferResult {
	ctx := request.Context
	input := request.Input
	requestId := util.GetRequestIDFromContext(ctx)
	log := util.GetLoggerFromContext(ctx).With(zap.String("request_id", requestId))

	s.mu.Lock()
	defer s.mu.Unlock()

	accountFrom, err := s.AccountRepo.GetById(input.AccountId)
	if err != nil {
		log.Error("error while retrieving account from database", zap.Error(err))
		return dto.NewTransferResult(nil, err)
	}

	log.Debug("retrieved account from database", zap.Any("account", accountFrom))

	accountTo, err := s.AccountRepo.GetById(input.TargetAccountId)
	if err != nil {
		log.Error("error while retrieving target account from database", zap.Error(err))
		return dto.NewTransferResult(nil, err)
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
			return dto.NewTransferResult(nil, err)
		}
		log.Debug("saved transfer", zap.Any("transfer", savedTransfer))
		transferOutput := output.NewCreateTransferOutput(savedTransfer.Id, savedTransfer.Success, savedTransfer.Date, savedTransfer.Reason)
		return dto.NewTransferResult(transferOutput, nil)
	}

	accountTo.Deposit(amount)
	log.Debug("deposit amount to target account", zap.Any("amount", amount))

	savedTransfer, err := s.TransferRepo.Save(*transfer)
	if err != nil {
		log.Error("error while saving transfer", zap.Error(err))
		return dto.NewTransferResult(nil, err)
	}

	log.Debug("saved transfer", zap.Any("transfer", savedTransfer))

	_, err = s.AccountRepo.Update(*accountFrom)

	if err != nil {
		log.Error("error while updating account", zap.Error(err))
		return dto.NewTransferResult(nil, err)
	}

	log.Debug("updated account", zap.Any("account", accountFrom))

	_, err = s.AccountRepo.Update(*accountTo)
	if err != nil {
		log.Error("error while updating target account", zap.Error(err))
		return dto.NewTransferResult(nil, err)
	}

	log.Debug("updated target account", zap.Any("target_account", accountTo))

	transferOutput := output.NewCreateTransferOutput(savedTransfer.Id, savedTransfer.Success, savedTransfer.Date, "")

	return dto.NewTransferResult(transferOutput, nil)
}
