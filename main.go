package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
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

	var currentImage = ""
	var firstRun = true
	for {
		if !firstRun {
			time.Sleep(time.Duration(conf.Preferences.Interval) * time.Second)
		}
		firstRun = false
		link := fetchLink(conf.Feed.Feed)
		if link.ID == 0 {
			debugPrint("[WallMaker] Failed to get link data")
			continue
		}
		if link.PostURL != currentImage {
			fmt.Println("[WallMaker] New image found", link.SetBy, ":", link.ID)
			setWallpaper(link.PostURL, link.Username, link.SetBy, link.UpdatedAt)
			currentImage = link.PostURL
		}
	}
}

func setWallpaper(url string, user string, setby string, timestamp time.Time) error {
	debugPrint("[WallMaker] Setting wallpaper to", url)
	if url == "" {
		return fmt.Errorf("No image URL found")
	}

	if conf.Preferences.SaveLocally {
		debugPrint("[WallMaker] Downloading image")
		ext := strings.Split(url, ".")[3]
		// download image to config.Preferences.SaveLocation
		timestamp := timestamp.Format("20060102150405")
		imageName := conf.Preferences.SaveLocation + "/" + timestamp + "_" + user + "_" + setby + "." + ext
		err := download(imageName, url)
		if err != nil {
			fmt.Println(err)
			return err
		}
		debugPrint("[WallMaker] Image downloaded")
	}

	if os.Getenv("TERMUX_VERSION") != "" {
		log.Println("[WARN] Setting wallpaper on Termux")
		log.Println("[WallMaker] Executing command:", conf.Termux.Cmd, conf.Termux.Args)
		//log.Println("[STAT] Executing command: termux-wallpaper -u", url)
		// replace {url} with url in args

		args := []string{}
		for _, arg := range conf.Termux.Args {
			args = append(args, strings.Replace(arg, "{URL}", url, -1))
		}
		cmd := exec.Command("termux-wallpaper", args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	} else {
		return wallpaper.SetFromURL(url)
	}
}

func debugPrint(msg ...any) {
	if conf.Base.Debug {
		log.Println(msg)
	}
}
