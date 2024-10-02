package main

import (
	"encoding/json"
	"fmt"
	"hidden_tunes/handler"
	"hidden_tunes/repository"
	"hidden_tunes/service"
	"io"
	"net/http"
	"net/url"
	"os"

	ht "hidden_tunes"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	version     = "v3.0"
	entity      = "tracks"
	maxListened = 30
)

func InitConfig() error {
	if err := godotenv.Load("configs/.env"); err != nil {
		return err
	}
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func gatherParams() url.Values {
	params := url.Values{}
	params.Add("client_id", os.Getenv("CLIENT_ID"))
	params.Add("format", "jsonpretty")
	params.Add("limit", "all")
	params.Add("order", "releasedate_desc")
	params.Add("include", "stats")
	params.Add("audioformat", "mp31")

	return params
}

func createRequest() *http.Request {
	secret := os.Getenv("CLIENT_SECRET")
	params := gatherParams().Encode()

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

func collectTunes(c *http.Client) {
	req := createRequest()
	resp, err := sendRequest(c, req)
	if err != nil {
		logrus.Println(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Println(err)
	}

	var APIResponse ht.APIResponse
	err = json.Unmarshal(body, &APIResponse)
	if err != nil {
		logrus.Println(err)
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
			AudioDownload:     res.AudioDownload,
			StatsRateListened: res.Stats.RateListenedTotal,
		}
		audios = append(audios, audio)
	}

	data, _ := json.Marshal(audios)

	err = os.WriteFile("response.json", data, 0o644)
	if err != nil {
		logrus.Fatalln(err)
	}
}

func main() {
	// logrus.SetFormatter(&logrus.JSONFormatter{})
	if err := InitConfig(); err != nil {
		logrus.Fatalln(err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		User:     viper.GetString("db.user"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalln(err)
	}

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	server := ht.NewServer(viper.GetString("port"), handler.InitRoutes())

	client := new(http.Client)
	collectTunes(client)

	if err := server.Run(); err != nil {
		logrus.Fatalln(err)
	}
}
