package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

const (
	version = "v3.0"
	entity
	subentity
)

func gatherParams() url.Values {
	params := url.Values{}
	return params

}

func createRequest() *http.Request {
	id := os.Getenv("CLIENT_ID")
	secret := os.Getenv("CLIENT_SECRET")
	params := gatherParams().Encode()

	url := fmt.Sprintf("https://api.jamendo.com/%s/%s/%s/?%s", version, entity, subentity, params)

	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Add("Authorization", bearer)
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

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if err := godotenv.Load(); err != nil {
		logrus.Fatalln(err)
	}

	client := &http.Client{}
	req := createRequest()
	resp, err := sendRequest(client, req)
	if err != nil {
		logrus.Println(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Println(err)
	}

	err = os.WriteFile("response.json", body, 0644)
	if err != nil {
		logrus.Fatalln(err)
	}

	fmt.Println("Response saved to response.json")
}
