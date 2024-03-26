package hello

import "go-be-template/internal/model/entity/wrapper"

func (u *Usecase) GetHelloMessageUsecase() (string, *wrapper.ErrorWrapper) {
	return u.helloRepo.GetHelloMessage()
}
