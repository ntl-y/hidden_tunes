package main

import (
	"fmt"
	"hidden_tunes/handler"
	"hidden_tunes/repository"
	"hidden_tunes/service"
	"os"

	ht "hidden_tunes"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig() error {
	if err := godotenv.Load("configs/.env"); err != nil {
		return err
	}
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	fmt.Println("Hello")
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
	} else {
		logrus.Println("Connected to db")
	}

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	go func() {
		err := service.FetchAudio()
		if err != nil {
			logrus.Error(err)
		}
	}()

	server := ht.NewServer(viper.GetString("port"), handler.InitRoutes())

	if err := server.Run(); err != nil {
		logrus.Fatalln(err)
	}
}
