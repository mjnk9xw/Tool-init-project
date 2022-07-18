package services

import (
	"{{PROJECT_NAME}}/repository"
)

// TODO: log, ctx, service, repository

type IService interface {
	{{I_USECASE}}
}

type Service struct {
	{{REPO_SERVICE}}
}

func New() *Service {
	return &Service{}
}

{{SET_REPO_SERVICE}}

func (s *Service) Build() IService {
	return s
}
