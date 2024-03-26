package hello

import "sinarmas/internal/model/entity/wrapper"

func (u *Usecase) GetHelloMessageUsecase() (string, *wrapper.ErrorWrapper) {
	return u.helloRepo.GetHelloMessage()
}
