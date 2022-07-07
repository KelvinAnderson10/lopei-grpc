package repository

import "lopei-grpc-server/model"

type LopeiRepository interface {
	RetrieveById(id int32) (model.Customer, error)
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

func NewLopeiRepository() LopeiRepository {
	repo := new(lopeiRepository)
	repo.db = []model.Customer{
		{LopeiId: 1, Balance: 5000},
		{LopeiId: 2, Balance: 1000},
		{LopeiId: 3, Balance: 15000},
	}
	return repo
}
