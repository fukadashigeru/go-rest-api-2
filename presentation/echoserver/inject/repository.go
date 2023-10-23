package inject

import (
	irepository "go-rest-api/domain/repository"
	"go-rest-api/repository"
)

func (injector *Injector) UserRepository() irepository.IUserRepository {
	return repository.NewUserRepository(injector.DB())
}

func (injector *Injector) TaskRepository() irepository.ITaskRepository {
	return repository.NewTaskRepository(injector.DB())
}
