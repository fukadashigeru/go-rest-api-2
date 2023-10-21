package inject

import (
	irepository "go-rest-api/interface/repository"
	"go-rest-api/repository"
)

func (i *Injector) UserRepository() irepository.IUserRepository {
	return repository.NewUserRepository(i.DB())
}
