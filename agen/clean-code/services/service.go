package services

import (
	"go.uber.org/zap"
	"{{PROJECT_NAME}}/repository"
	{{REDIS_LIB_V9}}
)

// TODO: log, ctx, service, repository

type IService interface {
	{{I_USECASE}}
}

type Service struct {
	log         *zap.Logger
	{{REDIS_CLIENT}}
	{{REPO_SERVICE}}
}

func New() *Service {
	return &Service{}
}

func (s *Service) SetLog(log *zap.Logger) *Service {
	s.log = log
	return s
}

{{SET_REDIS_SERVICE}}

{{SET_REPO_SERVICE}}

func (s *Service) Build() IService {
	return s
}
