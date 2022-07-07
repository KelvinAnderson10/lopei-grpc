package service

import (
	"context"
	"encoding/json"
	"lopei-grpc-server/repository"
)

type LopeiService struct {
	repo repository.LopeiRepository
	UnimplementedLopeiPaymentServer
}

func (l *LopeiService) CheckBalance(ctx context.Context, in *CheckBalanceMessage) (*ResultMessage, error) {
	lopeiId := in.LopeiId
	customer, err := l.repo.RetrieveById(lopeiId)
	if err != nil {
		return nil, err
	}

	c, err := json.Marshal(customer)
	if err != nil {
		return nil, err
	}
	resultMessage := &ResultMessage{
		Result: string(c),
		Error:  nil,
	}
	return resultMessage, nil
}

func (l *LopeiService) DoPayment(ctx context.Context, in *PaymentMessage) (*ResultMessage, error) {
	lopeiId := in.LopeiId
	amount := in.Amount
	customer, err := l.repo.RetrieveById(lopeiId)
	if err != nil {
		return nil, err
	}

	if customer.Balance < amount {
		return &ResultMessage{
			Result: "FAILED",
			Error: &Error{
				Code:    "X07",
				Message: "Insufficient Balance",
			},
		}, nil
	}
	err = l.repo.TransferBalance(lopeiId, in.LopeiReceiverId, amount)
	if err != nil {
		return &ResultMessage{
			Result: err.Error(),
			Error: &Error{
				Code:    "X08",
				Message: "Id not found",
			},
		}, nil
	}

	resultMessage := &ResultMessage{
		Result: "TRASNFER SUCCESS",
		Error:  nil,
	}
	return resultMessage, nil
}

func NewLopeiService(repo repository.LopeiRepository) *LopeiService {
	service := new(LopeiService)
	service.repo = repo
	return service
}
