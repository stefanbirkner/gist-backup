package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"

	"github.com/shurcooL/githubql"
	"golang.org/x/oauth2"
)

func main() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "0bf0f8891c6e1f9dbe9ef183ad926e41bcbe5d13"},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubql.NewClient(httpClient)

	type gist struct {
		Name        githubql.String
		Description githubql.String
	}
	var query struct {
		Viewer struct {
			Gists struct {
				Nodes []gist
			} `graphql:"gists(first: 100, privacy: ALL)"`
		}
	}
	err := client.Query(context.Background(), &query, nil)
	if err == nil {
		for _, element := range query.Viewer.Gists.Nodes {
			cloneGist(element.Name)
		}
	} else {
		log.Printf("Failed to read list of Gists: %v", err)
	}
}

func cloneGist(name githubql.String) {
	url := fmt.Sprintf("git@gist.github.com:%s.git", name)
	cmd := exec.Command("git", "clone", url)
	err := cmd.Run()
	if err != nil {
		log.Printf("Failed to clone gist %s: %v", name, err)
	}
}
