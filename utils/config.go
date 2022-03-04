package utils

import (
	"log"
	"os"
	"os/user"
)

var CONFIG_FILE_NAME string = "blackops"

func GetConfigFilePath() string {
	user, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	return user.HomeDir + "/." + CONFIG_FILE_NAME
}

func CreateConfigFile() {
	file_full_path := GetConfigFilePath()

	if _, err := os.Stat(file_full_path); os.IsNotExist(err) {
		file, _ := os.Create(file_full_path)
		defer file.Close()
	}
}
