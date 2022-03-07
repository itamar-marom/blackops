package models

import (
	"fmt"
	"os"

	git "github.com/go-git/go-git/v5"
	"github.com/itamar-marom/blackops/utils"
)

type Application struct {
	Name       string
	Repository string
	Properties string
}

func GetApplications(repository Repository) {
	r, _, err := utils.CloneRepositoryInMemory(repository.Token, repository.URL)
	utils.CheckError(err)

	w, err := r.Worktree()
	utils.CheckError(err)

	applications, err := w.Filesystem.ReadDir("applications")
	utils.CheckError(err)

	for _, file := range applications {
		println(file.Name())
	}
}

func GetApplication(appName string, repository Repository) {
	r, _, err := utils.CloneRepositoryInMemory(repository.Token, repository.URL)
	utils.CheckError(err)

	w, err := r.Worktree()
	utils.CheckError(err)

	applications, err := w.Filesystem.ReadDir("applications")
	utils.CheckError(err)

	isExists := false

	for _, file := range applications {
		if appName == file.Name() {
			println(file.Name())
			isExists = true
		}
	}

	if !isExists {
		fmt.Errorf("No such application " + appName + " in the given repository")
	}
}

func CreateApplication(appName string, repository Repository, propertiesFilePath string) error {
	r, fs, err := utils.CloneRepositoryInMemory(repository.Token, repository.URL)
	utils.CheckError(err)

	propertiesData, err := os.ReadFile(propertiesFilePath)
	utils.CheckError(err)

	appPath := "applications/" + appName

	fs.MkdirAll(appName, 0755)

	propertiesRepositoryPath := appPath + "/properties.yaml"
	propertiesFile, err := fs.Create(propertiesRepositoryPath)
	utils.CheckError(err)

	propertiesFile.Write(propertiesData)
	propertiesFile.Close()

	w, err := r.Worktree()
	utils.CheckError(err)

	w.Add(propertiesRepositoryPath)

	w.Commit("Added application: "+appName, &git.CommitOptions{})

	utils.PushRepositoryInMemory(repository.Token, r)

	fmt.Println("Created application: " + appName)
}
