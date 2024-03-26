package hello

import "sinarmas/internal/model/entity/wrapper"

func (r *Repository) GetHelloMessage() (string, *wrapper.ErrorWrapper) {
	return "Hello World!", nil
}
