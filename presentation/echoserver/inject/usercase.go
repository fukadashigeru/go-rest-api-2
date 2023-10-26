package inject

import (
	iusecase "go-rest-api/domain/usecase"
	"go-rest-api/usecase"
)

func (injector *Injector) UserUsecase() iusecase.IUserUsecase {
	return usecase.NewUserUsecase(injector.UserRepository(), injector.UserValidator())
}

func (injector *Injector) TaskUsecase() iusecase.ITaskUsecase {
	return usecase.NewTaskUsecase(injector.TaskRepository(), injector.TaskValidator())
}
