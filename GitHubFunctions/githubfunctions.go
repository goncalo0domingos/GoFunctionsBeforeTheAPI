package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v41/github"
	"golang.org/x/oauth2"
)

func main() {
	var my_token string = os.Getenv("GITHUB_TOKEN")
	list_all_repos(my_token)
	//create_singular_repo(my_token)
	//add_file_to_repo(my_token)
	//destroy_singular_repo(my_token)
	//list_N_pulls_for_repo(my_token)
	//list_all_branches(my_token)
	//make_new_branch(my_token, "general-test-branch") //"general-test-branch-2"
	//make_commit_change_to_branch(my_token, "general-test-branch")
	//make_pull_request_to_repo(my_token, "general-test-branch")
	//close_pull_request_from_repo(my_token, "test-repo-creation")

}

func list_all_repos(my_token string) { //working for auth user

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: my_token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		fmt.Println("Error on listing repositories")
	}

	for _, repo := range repos {
		fmt.Printf("Name: %s, Private: %v, URL: %s\n", *repo.Name, *repo.Private, *repo.HTMLURL)
	}
}

func destroy_singular_repo(my_token string) { //working

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: my_token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	repo_name := "test-repo-creation"
	owner := "goncalo0domingos"

	_, err := client.Repositories.Delete(ctx, owner, repo_name)
	if err != nil {
		fmt.Println("Error deleting %s repo", repo_name)
	} else {
		fmt.Println("Sucessfully deleted %s repo", repo_name)
	}
	list_all_repos(my_token)
}

func create_singular_repo(my_token string) { //working

	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: my_token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	repo := &github.Repository{
		Name:        github.String("test-repo-creation"),
		Private:     github.Bool(false),
		Description: github.String(""),
		HasIssues:   github.Bool(true),
		HasWiki:     github.Bool(true),
	}

	newRepo, _, err := client.Repositories.Create(ctx, "", repo)
	if err != nil {
		fmt.Println("Error creating repository: %v", err)
	}

	fmt.Println("New Repo created:\n Name: %s | HTML: %s", *newRepo.Name, *newRepo.HTMLURL)
}

func list_N_pulls_for_repo(my_token string) { //working

	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: my_token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	owner := "goncalo0domingos"
	repo_name := "test-repo-creation"

	options := &github.PullRequestListOptions{
		State: "open",
	}

	pullRequests, _, err := client.PullRequests.List(ctx, owner, repo_name, options)
	if err != nil {
		fmt.Printf("Error getting the information about number of open pull requests from repo: %s\n", repo_name)
	} else {
		fmt.Printf("Number of open Pull Requests for repo %s: %d\n", repo_name, len(pullRequests))
	}
}

func add_file_to_repo(my_token string) { //working

	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: my_token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	repo_name := "test-repo-creation"
	filePath := "NOT_A_README.txt"
	file_content := "NOT A DEFAULT README"
	owner := "goncalo0domingos"

	//encodedContent := base64.StdEncoding.EncodeToString([]byte(file_content))

	options := &github.RepositoryContentFileOptions{
		Message: github.String("Adding content with gitHub API - NOT A README.txt"),
		Content: []byte(file_content),
		Branch:  github.String("main"),
	}

	_, _, err := client.Repositories.CreateFile(ctx, owner, repo_name, filePath, options)
	if err != nil {
		fmt.Printf("Error adding file %s to main branch of respective repo: %s\n", filePath, repo_name)
		fmt.Println("Err Number: ", err)
	} else {
		fmt.Printf("File %s sucessfully created in repository %s\n", filePath, repo_name)
	}
}

func make_pull_request_to_repo(my_token string, name_of_branch string) { //
	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: my_token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	owner := "goncalo0domingos"
	repo_name := "test-repo-creation"

	newPR := &github.NewPullRequest{
		Title:               github.String("Pull Request from githubapi with go"),
		Head:                github.String(name_of_branch),
		Base:                github.String("main"),
		Body:                github.String("Trying the githubAPI with go"),
		MaintainerCanModify: github.Bool(true),
	}

	pr, _, err := client.PullRequests.Create(ctx, owner, repo_name, newPR)
	if err != nil {
		fmt.Printf("Error creating pull request: %v", err)
	} else {
		fmt.Printf("Pull request created: %s\n", pr.GetHTMLURL())
	}
}

func list_all_branches(my_token string) {
	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: my_token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	owner := "goncalo0domingos"
	repo := "test-repo-creation"

	branches, _, err := client.Repositories.ListBranches(ctx, owner, repo, nil)
	if err != nil {
		fmt.Printf("Error listing branches: %v", err)
	} else {
		for _, branch := range branches {
			fmt.Printf("Branch: %s\n", branch.GetName())
		}
	}
}

func make_new_branch(my_token string, name_of_branch string) { //working
	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: my_token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	owner := "goncalo0domingos"
	repo_name := "test-repo-creation"
	baseBranch := "main"
	newBranchName := name_of_branch

	ref, _, err := client.Git.GetRef(ctx, owner, repo_name, "refs/heads/"+baseBranch)
	if err != nil {
		fmt.Printf("Error getting reference of base branch: %v", err)
	}
	sha := ref.GetObject().GetSHA()
	//fmt.Printf("Base branch commit SHA: %s\n", sha)

	newRef := &github.Reference{
		Ref:    github.String("refs/heads/" + newBranchName),
		Object: &github.GitObject{SHA: github.String(sha)},
	}

	_, _, err = client.Git.CreateRef(ctx, owner, repo_name, newRef)
	if err != nil {
		fmt.Printf("Error creating new branch: %v", err)
	}

	fmt.Printf("Branch '%s' created successfully\n", newBranchName)
}

func make_commit_change_to_branch(my_token string, name_of_branch string) {

	ctx := context.Background()

	// Authenticate using the access token
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: my_token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	owner := "goncalo0domingos"
	repo := "test-repo-creation"
	branch := name_of_branch               // The branch you want to make changes to
	filePath := "path/to/NOT_A_README.txt" // The path to the file you want to change

	ref, _, err := client.Git.GetRef(ctx, owner, repo, "refs/heads/"+branch)
	if err != nil {
		fmt.Printf("Error getting reference of branch: %v", err)
	}
	sha := ref.GetObject().GetSHA()
	fmt.Printf("Branch latest commit SHA: %s\n", sha)

	commit, _, err := client.Git.GetCommit(ctx, owner, repo, sha)
	if err != nil {
		fmt.Printf("Error getting commit: %v", err)
	}
	treeSHA := commit.GetTree().GetSHA()

	content := []byte("New content from githubapi with go")
	newBlob := &github.Blob{
		Content:  github.String(string(content)),
		Encoding: github.String("utf-8"),
	}

	blob, _, err := client.Git.CreateBlob(ctx, owner, repo, newBlob)
	if err != nil {
		fmt.Printf("Error creating blob: %v", err)
	}

	entries := []*github.TreeEntry{
		{
			Path: github.String(filePath),
			Mode: github.String("100644"), // Standard file
			Type: github.String("blob"),
			SHA:  blob.SHA,
		},
	}
	newTree, _, err := client.Git.CreateTree(ctx, owner, repo, treeSHA, entries)
	if err != nil {
		fmt.Printf("Error creating tree: %v", err)
	}

	parent := []*github.Commit{commit}
	newCommit := &github.Commit{
		Message: github.String("New content added to NOT_A_README.txt"),
		Tree:    newTree,
		Parents: parent,
	}

	commitResponse, _, err := client.Git.CreateCommit(ctx, owner, repo, newCommit)
	if err != nil {
		fmt.Printf("Error creating commit: %v", err)
	}

	ref.Object.SHA = commitResponse.SHA
	_, _, err = client.Git.UpdateRef(ctx, owner, repo, ref, false)
	if err != nil {
		fmt.Printf("Error updating branch reference: %v", err)
	}

	fmt.Printf("Commit created successfully: %s\n", commitResponse.GetHTMLURL())
}

func close_pull_request_from_repo(my_token string, name_of_repo string) {
	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: my_token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	owner := "goncalo0domingos"
	repo_name := name_of_repo

	pullRequestNumber := 1 // neste caso apenas tenho o pull request number 1 open, por isso é que meto 1

	pr := &github.PullRequest{State: github.String("closed")}

	updatedPR, _, err := client.PullRequests.Edit(ctx, owner, repo_name, pullRequestNumber, pr)
	if err != nil {
		fmt.Printf("Error closing pull request: %v", err)
	}

	// Print the updated pull request status
	fmt.Printf("Closed Pull Request: %s (State: %s)\n", updatedPR.GetTitle(), updatedPR.GetState())
}
