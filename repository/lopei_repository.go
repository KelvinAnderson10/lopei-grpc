package repository

import (
	"errors"
	"lopei-grpc-server/model"
)

type LopeiRepository interface {
	RetrieveById(id int32) (model.Customer, error)
	TransferBalance(senderId int32, receiverId int32, amount float32) error
}

type lopeiRepository struct {
	db []model.Customer
}

func (l *lopeiRepository) RetrieveById(id int32) (model.Customer, error) {
	for _, customer := range l.db {
		if customer.LopeiId == id {
			return customer, nil
		}
	}
	return model.Customer{}, nil
}

func (l *lopeiRepository) TransferBalance(senderId int32, receiverId int32, amount float32) error {
	lengthData := int32(len(l.db)) + 1
	if senderId > lengthData || receiverId > lengthData {
		return errors.New("FAILED")
	}

	l.db[senderId-1].Balance = l.db[senderId-1].Balance - amount
	l.db[receiverId-1].Balance = l.db[receiverId-1].Balance + amount

	return nil
}

func NewLopeiRepository() LopeiRepository {
	repo := new(lopeiRepository)
	repo.db = []model.Customer{
		{LopeiId: 1, Balance: 5000},
		{LopeiId: 2, Balance: 1000},
		{LopeiId: 3, Balance: 15000},
	}
	return repo
}
