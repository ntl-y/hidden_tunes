package service

import "hidden_tunes/repository"

type AudioService struct {
	repo *repository.Repository
}

func NewAudioService(repo *repository.Repository) *AudioService {
	return &AudioService{repo: repo}
}
