package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shurcooL/githubv4"
	"github.com/sirupsen/logrus"
)

type repositoryInfoQuery struct {
	RateLimit  rateLimit
	Repository struct {
		DiskUsage  int
		ForkCount  int
		Stargazers struct {
			TotalCount int
		}
		Watchers struct {
			TotalCount int
		}
		IsPrivate  bool
		IsArchived bool
		IsDisabled bool
		IsFork     bool
		IsLocked   bool
		IsMirror   bool
		IsTemplate bool
		Languages  struct {
			Edges []struct {
				Size int
				Node struct {
					Name string
				}
			}
		} `graphql:"languages(first: 100)"`
	} `graphql:"repository(owner: $owner, name: $name)"`
}

type RepositoryInfo struct {
	// DiskUsage is returned in KBytes
	DiskUsage  int
	Forks      int
	Stargazers int
	Watchers   int
	IsPrivate  bool
	IsArchived bool
	IsDisabled bool
	IsFork     bool
	IsLocked   bool
	IsMirror   bool
	IsTemplate bool
	Languages  map[string]int
}

func (c *Client) RepositoryInfo(owner string, name string) (*RepositoryInfo, error) {
	variables := map[string]interface{}{
		"owner": githubv4.String(owner),
		"name":  githubv4.String(name),
	}

	var q repositoryInfoQuery

	err := c.client.Query(c.ctx, &q, variables)
	c.countRequest(owner, name, q.RateLimit)

	c.log.WithFields(logrus.Fields{
		"owner": owner,
		"name":  name,
		"cost":  q.RateLimit.Cost,
	}).Debugf("RepositoryInfo()")

	if err != nil {
		return nil, err
	}

	info := &RepositoryInfo{
		DiskUsage:  q.Repository.DiskUsage,
		Forks:      q.Repository.ForkCount,
		Stargazers: q.Repository.Stargazers.TotalCount,
		Watchers:   q.Repository.Watchers.TotalCount,
		IsPrivate:  q.Repository.IsPrivate,
		IsArchived: q.Repository.IsArchived,
		IsDisabled: q.Repository.IsDisabled,
		IsFork:     q.Repository.IsFork,
		IsLocked:   q.Repository.IsLocked,
		IsMirror:   q.Repository.IsMirror,
		IsTemplate: q.Repository.IsTemplate,
		Languages:  map[string]int{},
	}

	for _, lang := range q.Repository.Languages.Edges {
		info.Languages[lang.Node.Name] = lang.Size
	}

	return info, nil
}

type repositoriesNamesQuery struct {
	RateLimit       rateLimit
	RepositoryOwner struct {
		Repositories struct {
			Nodes []struct {
				Name string
			}
		} `graphql:"repositories(last: 100, isFork: false, isLocked: false, affiliations: OWNER)"`
	} `graphql:"repositoryOwner(login: $login)"`
}

func (c *Client) RepositoriesNames(login string) ([]string, error) {
	variables := map[string]interface{}{
		"login": githubv4.String(login),
	}

	var q repositoriesNamesQuery

	err := c.client.Query(c.ctx, &q, variables)

	if err != nil {
		c.log.Error(err)
		return nil, err
	}

	repos := []string{}
	for _, node := range q.RepositoryOwner.Repositories.Nodes {
		repos = append(repos, node.Name)
	}

	return repos, nil
}

type member struct {
	Login string `json:"login"`
}

const perPage = 100

func (c *Client) OrgMembers(org string) ([]string, error) {
	return c.orgMembersPage(org, 1)
}

func (c *Client) orgMembersPage(org string, page int) ([]string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/orgs/%s/members?per_page=%d&page=%d", org, perPage, page), nil)
	if err != nil {
		return nil, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("OrgMembers returned HTTP %d", res.StatusCode)
	}
	var resJSON []member
	if err := json.NewDecoder(res.Body).Decode(&resJSON); err != nil {
		return nil, err
	}
	members := make([]string, len(resJSON))
	for i := range members {
		members[i] = resJSON[i].Login
	}
	if len(members) == perPage {
		otherMembers, _ := c.orgMembersPage(org, page+1)
		members = append(members, otherMembers...)
	}
	return members, nil
}
