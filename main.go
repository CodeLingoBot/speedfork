package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage: speedfork <repo name> <git token>")
	}
	repo := os.Args[1]
	token := os.Args[2]

	ctx := context.Background()
	authedClient := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}))
	client := github.NewClient(authedClient)

	_, _, err := client.Repositories.CreateFork(ctx, "CodeLingoBot", repo, nil)
	if err != nil {
		if !strings.Contains(err.Error(), "job scheduled on GitHub side; try again later") {
			log.Fatal(err)
		}
	}
}
