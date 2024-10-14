package repository

import (
	ht "hidden_tunes"

	"github.com/jmoiron/sqlx"
)

type Audio interface {
	InsertAudioSlice(audios []ht.Audio) error
	GetRandomAudio() (ht.Audio, error)
}
type Repository struct {
	Audio
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Audio: NewAudioRepository(db)}
}
