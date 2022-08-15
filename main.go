package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/reujab/wallpaper"
)

func main() {
	err := setupConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(conf)

	var currentImage = -1

	for {
		link := fetchLink(conf.Feed.Feed)
		if link.ID != currentImage {
			fmt.Println("[WallMaker] New image found", link.ID)
			currentImage = link.ID
			if conf.Preferences.SaveLocally {
				fmt.Println("[WallMaker] Downloading image")
				ext := strings.Split(link.PostURL, ".")[3]

				timestamp := link.UpdatedAt.Format("20060102150405")
				imageName := timestamp + "_" + link.SetBy + "." + ext
				err := download(conf.Preferences.SaveLocation+"/"+imageName, link.PostURL)
				if err != nil {
					fmt.Println(err)
					return
				}

				err = setWallpaper(link.PostURL)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				continue
			}
			fmt.Println("[WallMaker] Setting wallpaper")
			err := setWallpaper(link.PostURL)
			if err != nil {
				fmt.Println(err)
			}
		}
		// convert into to time.Duration
		time.Sleep(time.Duration(conf.Preferences.Interval) * time.Second)
	}
}

func setWallpaper(url string) error {
	err := wallpaper.SetFromURL(url)
	if err != nil {
		return err
	}
	return nil
}
