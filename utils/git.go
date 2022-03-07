package utils

import (
	"github.com/go-git/go-billy/v5"
	memfs "github.com/go-git/go-billy/v5/memfs"
	git "github.com/go-git/go-git/v5"
	http "github.com/go-git/go-git/v5/plumbing/transport/http"
	memory "github.com/go-git/go-git/v5/storage/memory"
)

func authenticateRepositoryByBasic(repositoryToken string) *http.BasicAuth {
	// Authentication
	auth := &http.BasicAuth{
		Username: "_",
		Password: repositoryToken,
	}

	return auth
}

func CloneRepositoryInMemory(repositoryToken string, repositoryURL string) (*git.Repository, billy.Filesystem) {

	// Filesystem abstraction based on memory
	fs := memfs.New()
	// Git objects storer based on memory
	storer := memory.NewStorage()

	auth := authenticateRepositoryByBasic(repositoryToken)

	r, err := git.Clone(storer, fs, &git.CloneOptions{
		URL:  repositoryURL,
		Auth: auth,
	})
	CheckError(err)

	return r, fs
}

func PushRepositoryInMemory(repositoryToken string, r *git.Repository) {

	auth := authenticateRepositoryByBasic(repositoryToken)

	err := r.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth:       auth,
	})
	CheckError(err)
}
