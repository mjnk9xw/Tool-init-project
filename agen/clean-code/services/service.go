package services

// TODO: log, ctx, service, repository

type IService interface {
}

type service struct {
}

func New() IService {
	return &service{}
}
