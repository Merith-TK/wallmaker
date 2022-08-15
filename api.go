package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Link struct {
	ID               int         `json:"id"`
	Expires          interface{} `json:"expires"`
	Terms            string      `json:"terms"`
	Blacklist        string      `json:"blacklist"`
	PostURL          string      `json:"post_url"`
	PostThumbnailURL string      `json:"post_thumbnail_url"`
	PostDescription  string      `json:"post_description"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
	ResponseType     string      `json:"response_type"`
	ResponseText     string      `json:"response_text"`
	Username         string      `json:"username"`
	SetBy            string      `json:"set_by"`
	Online           bool        `json:"online"`
	URL              string      `json:"url"`
}

var (
	api = "https://walltaker.joi.how/api"
)

func fetchLink(id int) Link {
	var link Link
	getjson(api+"/links/"+fmt.Sprintf("%d", id)+".json", &link)
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
