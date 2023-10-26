package inject

import "go-rest-api/validator"

func (injector *Injector) UserValidator() validator.IUserValidator {
	return validator.NewUserValidator()
}

func (injector *Injector) TaskValidator() validator.ITaskValidator {
	return validator.NewTaskValidator()
}
