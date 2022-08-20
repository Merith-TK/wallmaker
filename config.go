package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/BurntSushi/toml"
)

var (
	conf       Config
	configfile = "config.toml"
)

type Config struct {
	Base struct {
		Base  string `toml:"base"`
		Debug bool   `toml:"debug"`
	} `toml:"base"`
	Feed struct {
		Feed int `toml:"feed"`
	} `toml:"feed"`
	Preferences struct {
		Interval     int    `toml:"interval"`
		Mode         string `toml:"mode"`
		SaveLocally  bool   `toml:"saveLocally"`
		SaveLocation string `toml:"savelocation"`
		Notification bool   `toml:"notification"`
	} `toml:"preferences"`
	Termux struct {
		Cmd  string   `toml:"cmd"`
		Args []string `toml:"arg"`
	} `toml:"termux"`
}

func setupConfig() error {
	if _, err := os.Stat(configfile); os.IsNotExist(err) {
		fmt.Println("[WallMaker] No config found, creating default config")
		f, err := os.Create(configfile)
		if err != nil {
			return err
		}

		// set default config
		conf.Base.Base = "https://walltaker.joi.how/api"
		conf.Feed.Feed = 0
		conf.Preferences.Interval = 10
		conf.Preferences.Mode = "normal"
		conf.Preferences.SaveLocally = false
		usr, _ := user.Current()
		conf.Preferences.SaveLocation = usr.HomeDir + "/Pictures/Wallpapers"
		conf.Preferences.Notification = false
		// write config to file
		toml.NewEncoder(f).Encode(conf)
		f.Close()

	} else {
		fmt.Println("[WallMaker] Found config, loading config")
		_, err := toml.DecodeFile(configfile, &conf)
		if err != nil {
			return err
		}

		usr, _ := user.Current()
		conf.Preferences.SaveLocation = strings.Replace(conf.Preferences.SaveLocation, "~", usr.HomeDir, 1)

	}
	// if save location doesn't exist, create it
	if _, err := os.Stat(conf.Preferences.SaveLocation); os.IsNotExist(err) {
		fmt.Println("[WallMaker] Save location doesn't exist, creating")
		os.MkdirAll(conf.Preferences.SaveLocation, 0755)
	}
	return nil
}
