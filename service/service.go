package service

import "hidden_tunes/repository"

type Audio interface {
}

type Service struct {
	Audio
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Audio: NewAudioService(repo)}
}
