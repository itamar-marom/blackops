package utils

import (
	"github.com/go-git/go-billy/v5"
	memfs "github.com/go-git/go-billy/v5/memfs"
	git "github.com/go-git/go-git/v5"
	http "github.com/go-git/go-git/v5/plumbing/transport/http"
	memory "github.com/go-git/go-git/v5/storage/memory"
)

func CloneRepositoryInMemory(password string, repository_url string) (*git.Repository, billy.Filesystem, error) {

	// Filesystem abstraction based on memory
	fs := memfs.New()
	// Git objects storer based on memory
	storer := memory.NewStorage()

	// Authentication
	auth := &http.BasicAuth{
		Username: "_",
		Password: password,
		// Password: "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",
	}

	repository := repository_url

	r, err := git.Clone(storer, fs, &git.CloneOptions{
		URL:  repository,
		Auth: auth,
	})

	if err != nil {
		return nil, nil, err
	}

	return r, fs, nil
}

func PushRepositoryInMemory(password string, r *git.Repository) {

	// Authentication
	auth := &http.BasicAuth{
		Username: "_",
		Password: password,
		// Password: "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",
	}

	err := r.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth:       auth,
	})

	CheckError(err)
}
