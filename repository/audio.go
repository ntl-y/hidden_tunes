package repository

import (
	"fmt"
	ht "hidden_tunes"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type AudioRepository struct {
	db *sqlx.DB
}

func NewAudioRepository(db *sqlx.DB) *AudioRepository {
	return &AudioRepository{db: db}
}

func (r *AudioRepository) insertAudio(audio ht.Audio) error {
	query := fmt.Sprintf("INSERT INTO %s (jamendo_id, name, artist_name, album_name, album_image, audio, audiodownload, rate_listened_total) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) ON CONFLICT (jamendo_id) DO NOTHING", audioTable)
	args := []interface{}{audio.ID, audio.Name, audio.ArtistName, audio.AlbumName, audio.AlbumImage, audio.Audio, audio.AudioDownload, audio.StatsRateListened}
	_, err := r.db.Exec(query, args...)
	return err
}

func (r *AudioRepository) InsertAudioSlice(audios []ht.Audio) error {

	for _, a := range audios {
		if err := r.insertAudio(a); err != nil {
			return err
		}
	}
	logrus.Info(time.Now().Format("01/02/2006 15:04:05"), ": New Audio Inserted")
	return nil
}
