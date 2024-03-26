package app

import "sinarmas/internal/model/config"

type CommonResource struct {
}

func NewCommonResource(cfg *config.Config) (*CommonResource, error) {
	return &CommonResource{}, nil
}
