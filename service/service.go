package service

import (
	ht "hidden_tunes"
	"hidden_tunes/repository"
)

type Audio interface {
	FetchAudio() error
	GetRandomAudio() (ht.Audio, error)
}

type Service struct {
	Audio
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Audio: NewAudioService(repo)}
}
