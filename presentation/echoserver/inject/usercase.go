package inject

import (
	iusecase "go-rest-api/interface/usecase"
	"go-rest-api/usecase"
)

func (injector *Injector) UserUsecase() iusecase.IUserUsecase {
	return usecase.NewUserUsecase(injector.UserRepository())
}
