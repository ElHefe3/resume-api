package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	NextcloudURL      		string
	NextcloudUsername 		string
	NextcloudPassword 		string
	NextcloudFilesDirectory string
	AuthToken 				string
	FeUrl					string
}

var Cfg Config

func Load() {
	_ = godotenv.Load()

	Cfg = Config{
		NextcloudURL:      			getEnv("NEXTCLOUD_URL", ""),
		NextcloudUsername: 			getEnv("NEXTCLOUD_USERNAME", ""),
		NextcloudPassword: 			getEnv("NEXTCLOUD_PASSWORD", ""),
		NextcloudFilesDirectory: 	getEnv("NEXTCLOUD_FILES_DIRECTORY", ""),
		AuthToken:					getEnv("AUTH_TOKEN", ""),
		FeUrl: 						getEnv("FE_URL", ""),
	}

	if Cfg.NextcloudURL == "" || Cfg.NextcloudUsername == "" || Cfg.NextcloudPassword == "" || Cfg.NextcloudFilesDirectory == "" || Cfg.AuthToken == "" {
		log.Fatal("Missing required environment variables")
	}
}

func getEnv(key, defaultValue string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return defaultValue
}
