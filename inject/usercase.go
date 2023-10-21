package inject

import (
	iusecase "go-rest-api/interface/usecase"
	"go-rest-api/usecase"
)

func (i *Injector) UserUsecase() iusecase.IUserUsecase {
	return usecase.NewUserUsecase(i.UserRepository())
}
