package service

import (
	"encoding/json"
	"errors"
	"fmt"
	ht "hidden_tunes"
	"hidden_tunes/repository"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

const (
	version     = "v3.0"
	entity      = "tracks"
	maxListened = 1000
	limit       = 200
)

type AudioService struct {
	repo   *repository.Repository
	client *http.Client
	offset int
}

func NewAudioService(repo *repository.Repository) *AudioService {
	return &AudioService{
		repo:   repo,
		client: new(http.Client),
		offset: 0,
	}
}

func (s *AudioService) GetRandomAudio() (ht.Audio, error) {
	return s.repo.GetRandomAudio()
}

func gatherParams(offset int) url.Values {
	params := url.Values{}
	params.Add("client_id", os.Getenv("CLIENT_ID"))
	params.Add("format", "jsonpretty")
	params.Add("limit", "all")
	params.Add("offset", strconv.Itoa(offset))
	params.Add("order", "releasedate_desc")
	params.Add("include", "stats")
	params.Add("audioformat", "mp31")

	return params
}

func createRequest(offset int) *http.Request {
	secret := os.Getenv("CLIENT_SECRET")
	params := gatherParams(offset).Encode()

	url := fmt.Sprintf("https://api.jamendo.com/%s/%s/?%s", version, entity, params)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", secret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req
}

func sendRequest(client *http.Client, req *http.Request) (*http.Response, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *AudioService) collectAudio() error {
	req := createRequest(s.offset)
	resp, err := sendRequest(s.client, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var APIResponse ht.APIResponse
	err = json.Unmarshal(body, &APIResponse)
	if err != nil {
		return err
	}

	if APIResponse.Headers.Status == "failed" {
		return errors.New(APIResponse.Headers.ErrorMessage)
	}

	var audios []ht.Audio
	for _, res := range APIResponse.Results {
		if !res.AudioDownloadAllowed || res.Stats.RateListenedTotal > maxListened {
			continue
		}
		audio := ht.Audio{
			ID:                res.ID,
			Name:              res.Name,
			ArtistName:        res.ArtistName,
			AlbumName:         res.AlbumName,
			Audio:             res.AlbumName,
			AlbumImage:        res.AlbumImage,
			AudioDownload:     res.AudioDownload,
			StatsRateListened: res.Stats.RateListenedTotal,
		}
		audios = append(audios, audio)
	}
	return s.repo.InsertAudioSlice(audios)
}

func (s *AudioService) FetchAudio() error {
	for {
		err := s.collectAudio()
		if err != nil {
			return err
		}
		s.offset += limit
		time.Sleep(30 * time.Minute)
	}
}

func (s *AudioService) DBValidate() {
}
