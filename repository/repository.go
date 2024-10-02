package repository

import "github.com/jmoiron/sqlx"

type Audio interface {
}
type Repository struct {
	Audio
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Audio: NewAudioRepository(db)}
}
