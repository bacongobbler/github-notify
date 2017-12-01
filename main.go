package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const (
	EnvGithubAuthToken   = "GITHUB_AUTH_TOKEN"
	EnvGithubOwner       = "GITHUB_OWNER"
	EnvGithubRepo        = "GITHUB_REPO"
	EnvGithubRef         = "GITHUB_REF"
	EnvGithubState       = "GITHUB_STATE"
	EnvGithubTargetURL   = "GITHUB_TARGET_URL"
	EnvGithubDescription = "GITHUB_DESCRIPTION"
	EnvGithubContext     = "GITHUB_CONTEXT"
)

// SetRepoStatus sets the status on a particular commit in a repo.
func setRepoStatus(client *github.Client, repo string, ref string, status *github.RepoStatus) error {
	parts := strings.SplitN(repo, "/", 3)
	if len(parts) != 3 {
		return fmt.Errorf("repo name %q is malformed", repo)
	}
	c := context.Background()
	_, _, err := client.Repositories.CreateStatus(
		c,
		parts[1],
		parts[2],
		ref,
		status)
	return err
}

func main() {
	authToken := os.Getenv(EnvGithubAuthToken)
	if authToken == "" {
		fmt.Fprintln(os.Stderr, "$GITHUB_AUTH_TOKEN is not set.")
		os.Exit(1)
	}

	owner := os.Getenv(EnvGithubOwner)
	if owner == "" {
		fmt.Fprintln(os.Stderr, "$GITHUB_OWNER is not set.")
		os.Exit(1)
	}

	repo := os.Getenv(EnvGithubRepo)
	if repo == "" {
		fmt.Fprintln(os.Stderr, "$GITHUB_REPO is not set.")
		os.Exit(1)
	}

	ref := os.Getenv(EnvGithubRef)
	if ref == "" {
		fmt.Fprintln(os.Stderr, "$GITHUB_REF is not set.")
		os.Exit(1)
	}

	state := os.Getenv(EnvGithubState)
	if state == "" {
		fmt.Fprintln(os.Stderr, "$GITHUB_STATE is not set.")
		os.Exit(1)
	}
	targetURL := os.Getenv(EnvGithubTargetURL)
	description := os.Getenv(EnvGithubDescription)
	ctx := os.Getenv(EnvGithubContext)

	t := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: authToken})
	tc := oauth2.NewClient(context.Background(), t)
	ghClient := github.NewClient(tc)

	status := &github.RepoStatus{
		State:       &state,
		TargetURL:   &targetURL,
		Description: &description,
		Context:     &ctx,
	}

	_, resp, err := ghClient.Repositories.CreateStatus(context.Background(), owner, repo, ref, status)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating status: %v\n", err)
		os.Exit(2)
	}

	fmt.Printf("GitHub responded with status code %d\n", resp.StatusCode)
}
