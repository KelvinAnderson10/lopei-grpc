package manager

import "lopei-grpc-server/repository"

type RepositoryManager interface {
	LopeiRepository() repository.LopeiRepository
}

type repositoryManager struct {
	lopeiRepository repository.LopeiRepository
}

func (r *repositoryManager) LopeiRepository() repository.LopeiRepository {
	return r.lopeiRepository
}

func NewRepositoryManager() RepositoryManager {
	repo := new(repositoryManager)
	repo.lopeiRepository = repository.NewLopeiRepository()
	return repo
}
