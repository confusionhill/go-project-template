package hello

import (
	"sinarmas/internal/model/config"
	"sinarmas/internal/repository/hello"
)

type Usecase struct {
	cfg       *config.Config
	helloRepo *hello.Repository
}

func New(cfg *config.Config, helloRepo *hello.Repository) (*Usecase, error) {
	return &Usecase{
		cfg:       cfg,
		helloRepo: helloRepo,
	}, nil
}
