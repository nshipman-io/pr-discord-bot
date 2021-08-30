package github

import (
	"context"
	"fmt"
	"github.com/nshipman-io/pr-discord-bot/config"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type Repo struct {
	Name string
	Owner string
	PullRequests []PullRequest
}
type PullRequest struct {
	Number int
}

func GetOpenPrs() ([]Repo) {
	var Repos []Repo
	fmt.Println(config.GitToken)
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken : config.GitToken},
	)

	httpClient := oauth2.NewClient(context.Background(), src)

	c := githubv4.NewClient(httpClient)

	var q struct {
		Repository struct {
			PullRequests struct {
				Nodes []PullRequest
			}`graphql:"pullRequests(last: 5, states: OPEN)"`
		} `graphql:"repository(owner: $owner, name: $name)"`
	}

	for _,v := range config.Repos {
		variables := map[string]interface{}{
			"owner": githubv4.String(v.Owner),
			"name": githubv4.String(v.Name),
		}


		err := c.Query(context.Background(), &q, variables)
		if err != nil {
			fmt.Println(err.Error())
		}
		r := Repo{
			Name: v.Name,
			Owner: v.Owner,
			PullRequests: q.Repository.PullRequests.Nodes,
		}
		Repos = append(Repos, r)
	}
	return Repos
}