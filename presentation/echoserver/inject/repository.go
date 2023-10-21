package inject

import (
	irepository "go-rest-api/interface/repository"
	"go-rest-api/repository"
)

func (injector *Injector) UserRepository() irepository.IUserRepository {
	return repository.NewUserRepository(injector.DB())
}
