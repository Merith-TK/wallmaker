package main

// import (
// 	"fmt"
// 	"log"
// 	"strings"
// 	"time"
// )

// func disabled() {
// 	var currentImage = ""
// 	for {
// 		link := fetchLink(conf.Feed.Feed)
// 		if link.ID == 0 {
// 			debugPrint("[WallMaker] Failed to get link data")
// 			continue
// 		}
// 		if link.PostURL != currentImage {
// 			fmt.Println("[WallMaker] New image found", link.SetBy, ":", link.ID)
// 			// TODO: push notification
// 			//notify("New image found", link.SetBy+": "+link.URL)
// 			currentImage = link.PostURL
// 			if conf.Preferences.SaveLocally {
// 				debugPrint("[WallMaker] Downloading image")
// 				ext := strings.Split(link.PostURL, ".")[3]

// 				timestamp := link.UpdatedAt.Format("20060102150405")
// 				imageName := timestamp + "_" + link.SetBy + "." + ext
// 				err := download(conf.Preferences.SaveLocation+"/"+imageName, link.PostURL)
// 				if err != nil {
// 					fmt.Println(err)
// 					return
// 				}
// 				debugPrint("[WallMaker] Image downloaded")
// 			}
// 			fmt.Println("[WallMaker] Setting wallpaper")
// 			err := setWallpaper(link.PostURL)
// 			if err != nil {
// 				log.Fatalln("Failed to set wallpaper", err)
// 			}
// 		}
// 		time.Sleep(time.Duration(conf.Preferences.Interval) * time.Second)
// 	}
// }
