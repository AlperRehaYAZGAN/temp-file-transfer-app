package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

type config struct {
	App struct {
		Env        string `mapstructure:"env" default:"dev"`
		Port       string `mapstructure:"port" default:"9090"`
		Version    string `mapstructure:"version" default:"1.0.0"`
		UploadsDir string `mapstructure:"uploads_dir" default:"uploads"`
	}
}

var C config
var Pwd string

func ReadConfig(processCwdir string) {
	Config := &C

	// if .env.yaml exist read it
	if _, err := os.Stat(processCwdir + "/config" + "/.env.yaml"); err == nil {
		viper.SetConfigName(".env")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(filepath.Join(processCwdir, "config"))
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}

		if err := viper.Unmarshal(&Config); err != nil {
			log.Fatalf("Error unable to decode into struct, %v", err)
			os.Exit(1)
		}

		spew.Dump(C)
	} else {
		// if .env.yaml not exist set default values
		Config.App.Env = "dev"
		Config.App.Port = "9090"
		Config.App.Version = "1.0.0"
		Config.App.UploadsDir = "uploads"
	}

	// if ../uploads not exist create it
	if _, err := os.Stat(processCwdir + "/" + Config.App.UploadsDir); os.IsNotExist(err) {
		err = os.Mkdir(processCwdir+"/"+Config.App.UploadsDir, 0777)
		if err != nil {
			log.Fatal("Error creating uploads directory, ", err)
		}
	}

}
