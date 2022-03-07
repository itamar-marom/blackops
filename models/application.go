package models

import (
	"bufio"
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

func getApplicationProperties(filePath string, w *git.Worktree) string {
	file, err := w.Filesystem.Open(filePath)
	utils.CheckError(err)

	scanner := bufio.NewScanner(file)

	properties := ""

	for scanner.Scan() {
		properties = properties + "\n" + scanner.Text()
	}

	err = scanner.Err()
	utils.CheckError(err)

	return properties
}

func GetApplications(repository Repository) []Application {
	r, _ := utils.CloneRepositoryInMemory(repository.Token, repository.URL)

	w, err := r.Worktree()
	utils.CheckError(err)

	applicationsDir := "applications"

	applicationsDirectories, err := w.Filesystem.ReadDir(applicationsDir)
	utils.CheckError(err)

	applications := []Application{}

	for _, file := range applicationsDirectories {
		appPath := applicationsDir + "/" + file.Name() + "/properties.yaml"
		app := new(Application)
		app.Name = file.Name()
		app.Repository = repository.Name
		app.Properties = getApplicationProperties(appPath, w)
		applications = append(applications, *app)
	}

	return applications
}

func GetApplication(appName string, repository Repository) (Application, int, error) {
	applications := GetApplications(repository)

	for index, app := range applications {
		if appName == app.Name {
			return app, index, nil
		}
	}

	return Application{}, 0, fmt.Errorf("Application does not exists")
}

func PrintApplication(app Application) {
	fmt.Println("Name: " + app.Name)
	fmt.Println("Repository: " + app.Repository)
	fmt.Println("Properties:")
	fmt.Println("----------" + app.Properties)
	fmt.Println("==============================")
}

func PrintApplicationByName(appName string, repo Repository) {
	app, _, err := GetApplication(appName, repo)
	utils.CheckError(err)

	PrintApplication(app)
}

func PrintApplications(repo Repository) {
	applications := GetApplications(repo)

	for index, app := range applications {
		fmt.Print(index + 1)
		fmt.Println(": ")
		PrintApplication(app)
	}
}

func CreateApplication(appName string, repository Repository, propertiesFilePath string) {
	r, fs := utils.CloneRepositoryInMemory(repository.Token, repository.URL)

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

func DeleteApplication(appName string, repository Repository) {
	r, fs := utils.CloneRepositoryInMemory(repository.Token, repository.URL)

	appPath := "applications/" + appName

	fs.Remove(appPath)

	w, err := r.Worktree()
	utils.CheckError(err)

	w.Remove(appPath)

	w.Commit("Deleted application: "+appName, &git.CommitOptions{})

	utils.PushRepositoryInMemory(repository.Token, r)

	fmt.Println("Deleted application: " + appName)
}
