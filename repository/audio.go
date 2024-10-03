package repository

import (
	"fmt"
	ht "hidden_tunes"
	"os"
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

func (r *AudioRepository) CleanAudio() error {
	for {
		err := os.RemoveAll("./web/music/*")
		if err != nil {
			return err
		}
		logrus.Infof("%s: AudioDir Cleaned", time.Now().Format("01/02/2006 15:04:05"))
		time.Sleep(5 * time.Hour)
	}

}

func (r *AudioRepository) GetRandomAudio() (audio ht.Audio, err error) {
	query := fmt.Sprintf("select jamendo_id, name, artist_name, album_name, album_image, audio, audiodownload, rate_listened_total from %s order by random() limit 1", audioTable)
	err = r.db.Get(&audio, query)
	return audio, err
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
	logrus.Infof("%s: %d Audio Inserted", time.Now().Format("01/02/2006 15:04:05"), len(audios))
	return nil
}
