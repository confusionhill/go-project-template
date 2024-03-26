package app

import (
	"sinarmas/internal/model/config"
	"sinarmas/internal/repository/hello"
)

type CommonRepository struct {
	helloRepo *hello.Repository
}

func NewCommonRepository(cfg *config.Config, rsc *CommonResource) (*CommonRepository, error) {
	helloRepo, err := hello.New(cfg)
	if err != nil {
		return nil, err
	}
	return &CommonRepository{
		helloRepo: helloRepo,
	}, nil
}
