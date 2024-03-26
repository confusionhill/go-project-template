package hello

import "go-be-template/internal/model/entity/wrapper"

func (r *Repository) GetHelloMessage() (string, *wrapper.ErrorWrapper) {
	return "Hello World!", nil
}
