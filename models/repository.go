package models

import (
	"fmt"
	"io/ioutil"

	"github.com/itamar-marom/blackops/utils"
	"gopkg.in/yaml.v2"
)

type Repository struct {
	Name  string `yaml:"name"`
	URL   string `yaml:"url"`
	Token string `yaml:"token"`
}

func GetConfig() map[string][]Repository {
	config := make(map[string][]Repository)

	config_file_path := utils.GetConfigFilePath()
	yaml_config, err := ioutil.ReadFile(config_file_path)

	if err != nil {
		println(err.Error())
	}

	err = yaml.Unmarshal(yaml_config, &config)
	if err != nil {
		println(err.Error())
	}

	return config
}

func GetRepository(name string) (Repository, error) {
	config := GetConfig()
	repositories := config["repositories"]

	for _, repository := range repositories {
		if name == repository.Name {
			return repository, nil
		}
	}

	return Repository{}, fmt.Errorf("Repository does not exists")
}

func AddNewRepository(newRepository Repository) {
	config_file_path := utils.GetConfigFilePath()
	config := GetConfig()

	repo, _ := GetRepository(newRepository.Name)
	if repo.Name == newRepository.Name {
		println("Repository already exists")
	} else {
		config["repositories"] = append(config["repositories"], newRepository)

		newConfig, err := yaml.Marshal(&config)
		if err != nil {
			println(err.Error())
		}

		err = ioutil.WriteFile(config_file_path, newConfig, 0)

		if err != nil {
			println(err.Error())
		} else {
			println("Repository " + newRepository.Name + " has been added")
		}
	}
}
