package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v48/github"
)

func main() {
	fmt.Println("Works")

	cli := github.NewClient(nil)
	repo, _, err := cli.Repositories.Get(context.Background(), "google", "go-github")
	if err != nil {
		panic(err)
	}

	fmt.Println(*repo.ID)
	fmt.Println(*repo.Name)
	fmt.Println(*repo.FullName)
	fmt.Println(*repo.HTMLURL)
	fmt.Println(*repo.Archived)
	fmt.Println(*repo.Description)
	fmt.Println(*repo.ForksCount)
	fmt.Println(*repo.StargazersCount)
	fmt.Println(repo.Topics)
	fmt.Println(*repo.PushedAt)
}
