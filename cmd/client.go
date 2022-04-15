package cmd

import (
	"context"

	"github.com/demachitse/manage-github-labels/config"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func GithubClient(ctx context.Context) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: *config.Data.Token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return client
}

