package repository

import "github.com/jmoiron/sqlx"

type AudioRepository struct {
	db *sqlx.DB
}

func NewAudioRepository(db *sqlx.DB) *AudioRepository {
	return &AudioRepository{db: db}
}
