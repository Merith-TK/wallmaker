package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
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
			setWallpaper(link)
			currentImage = link.PostURL
		}
	}
}

func setWallpaper(link Link) error {
	debugPrint("[WallMaker] Setting wallpaper to", link.PostURL)
	if link.PostURL == "" {
		return fmt.Errorf("No image URL found")
	}
	url := link.PostURL
	if conf.Preferences.SaveLocally {
		debugPrint("[WallMaker] Downloading image")
		ext := strings.Split(url, ".")[3]
		timestamp := link.UpdatedAt.Format("20060102150405")
		feed := fmt.Sprintf("%d", conf.Feed.Feed)

		setby := ""
		if link.SetBy != "" {
			setby = link.SetBy
		} else {
			setby = "unknown"
		}

		filename := fmt.Sprintf("%s_%s_%s_%s.%s", timestamp, feed, link.Username, setby, ext)
		imageName := conf.Preferences.SaveLocation + "/" + filename
		err := download(imageName, url)
		if err != nil {
			fmt.Println(err)
			return err
		}
		debugPrint("[WallMaker] Image downloaded")
	}

	// if /data/data/com.termux/ exists, use termux-wallpaper
	if _, err := os.Stat("/data/data/com.termux/files/home/"); err == nil {
		log.Println("[WARN] Setting wallpaper on Termux")
		log.Println("[WallMaker] Executing command:", conf.Termux.Cmd, conf.Termux.Args)
		//log.Println("[STAT] Executing command: termux-wallpaper -u", url)
		// replace {url} with url in args
		// args := []string{}
		// for _, arg := range conf.Termux.Args {
		// 	args = append(args, strings.Replace(arg, "{URL}", url, -1))
		// }
		cmd := exec.Command("termux-wallpaper", "-u", url)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		ret := cmd.Run()
		if ret != nil {
			return ret
		}
		sendNotification("[Wallmaker]", "Someone updated your wallpaper", true)
	} else {
		sendNotification("[Wallmaker]", "Someone updated your wallpaper", false)
		return wallpaper.SetFromURL(url)
	}
	return nil
}

func sendNotification(title string, body string, termux bool) {
	if runtime.GOOS == "windows" {
		return
	}
	if termux {
		cmd := exec.Command("termux-notification", "-t", title, "-b", body)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		ret := cmd.Run()
		if ret != nil {
			return
		}
	} else {
		cmd := exec.Command("notify-send", title, body)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		ret := cmd.Run()
		if ret != nil {
			return
		}
	}
}

func debugPrint(msg ...any) {
	if conf.Base.Debug {
		log.Println(msg)
	}
}
