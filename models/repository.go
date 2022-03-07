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

func getRepositoriesConfig() map[string][]Repository {
	config := make(map[string][]Repository)

	config_file_path := utils.GetConfigFilePath()
	yaml_config, err := ioutil.ReadFile(config_file_path)
	utils.CheckError(err)

	err = yaml.Unmarshal(yaml_config, &config)
	utils.CheckError(err)

	return config
}

func getRepositories() []Repository {
	config := getRepositoriesConfig()
	return config["repositories"]
}

func GetRepository(name string) (Repository, int, error) {
	repositories := getRepositories()

	for index, repository := range repositories {
		if name == repository.Name {
			return repository, index, nil
		}
	}

	return Repository{}, 0, fmt.Errorf("Repository does not exists")
}

func PrintRepository(repo Repository) {
	fmt.Println("Name: " + repo.Name)
	fmt.Println("URL: " + repo.URL)
}

func PrintRepositories() {
	repositories := getRepositories()

	for _, repo := range repositories {
		PrintRepository(repo)
	}
}

func CreateRepository(newRepository Repository) {
	config_file_path := utils.GetConfigFilePath()
	config := getRepositoriesConfig()

	repo, _, _ := GetRepository(newRepository.Name)
	if repo.Name == newRepository.Name {
		println("Repository already exists")
	} else {
		config["repositories"] = append(config["repositories"], newRepository)

		newConfig, err := yaml.Marshal(&config)
		utils.CheckError(err)

		err = ioutil.WriteFile(config_file_path, newConfig, 0)
		utils.CheckError(err)

		println("Repository " + newRepository.Name + " has been added")
	}
}

func DeleteRepository(repositoryName string) {
	_, index, err := GetRepository(repositoryName)
	utils.CheckError(err)

	config_file_path := utils.GetConfigFilePath()
	config := getRepositoriesConfig()
	repositories := config["repositories"]
	config["repositories"] = append(repositories[:index], repositories[index+1:]...)

	newConfig, err := yaml.Marshal(&config)
	utils.CheckError(err)

	err = ioutil.WriteFile(config_file_path, newConfig, 0)
	utils.CheckError(err)

	println("Repository " + repositoryName + " has been deleted")
}
