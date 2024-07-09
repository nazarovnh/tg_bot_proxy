package service

import "tg_bot_proxy/internal/repository"


type Service struct {
}

// function to initialize all services
func Init(repos *repository.Repository) *Service {
	return &Service{
	}
}