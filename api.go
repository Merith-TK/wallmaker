package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Link struct {
	ID               int       `json:"id"`
	Expires          time.Time `json:"expires"`
	Username         string    `json:"username"`
	Terms            string    `json:"terms"`
	Blacklist        string    `json:"blacklist"`
	PostURL          string    `json:"post_url"`
	PostThumbnailURL string    `json:"post_thumbnail_url"`
	PostDescription  string    `json:"post_description"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	SetBy            string    `json:"set_by"`
	ResponseType     string    `json:"response_type"`
	ResponseText     string    `json:"response_text"`
	Online           bool      `json:"online"`
}

var (
	api = "https://walltaker.joi.how/api"
)

func fetchLink(id int) Link {
	var link Link
	// convert id to string
	fetch := fmt.Sprintf("%s/links/%d.json", api, id)
	debugPrint("[WallMaker] Fetching link", fetch)
	getjson(fetch, &link)
	return link
}

func getjson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	if r.StatusCode != 200 {
		return errors.New("received non 200 response code")
	}
	defer r.Body.Close()
	log.Println("[WallMaker] Parsing json", url)
	log.Println("[WallMaker] Parsing json", r.Body)
	return json.NewDecoder(r.Body).Decode(target)
}

func download(fileName string, URL string) error {
	//Get the response bytes from the url
	fmt.Println(fileName, "\n", URL)
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		response, err := http.Get(URL)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		if response.StatusCode != 200 {
			return errors.New("received non 200 response code")
		}
		//Create a empty file
		file, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer file.Close()

		//Write the bytes to the fiel
		_, err = io.Copy(file, response.Body)
		if err != nil {
			return err
		}
	}
	return nil
}
