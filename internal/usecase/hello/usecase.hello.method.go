package hello

import (
	"go-be-template/internal/model/entity"
	"go-be-template/internal/model/entity/wrapper"
)

func (u *Usecase) GetHelloMessageUsecase() (*entity.HelloEntity, *wrapper.ErrorWrapper) {
	return u.helloRepo.GetHelloMessage()
}
