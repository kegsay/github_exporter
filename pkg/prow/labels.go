package prow

import (
	"fmt"
	"regexp"
	"strings"

	"go.xrstf.de/github_exporter/pkg/github"
)

func PullRequestLabelNames() []string {
	return []string{"size", "T", "priority"}
}

func PullRequestLabels(pr *github.PullRequest) []string {
	return []string{
		prefixedLabel("size", pr.Labels),
		prefixedLabel("T", pr.Labels),
		prefixedLabel("priority", pr.Labels),
	}
}

func IssueLabelNames() []string {
	return []string{"T"}
}

func IssueLabels(issue *github.Issue) []string {
	return []string{
		prefixedLabel("T", issue.Labels),
	}
}

func prefixedLabel(prefix string, labels []string) string {
	prefix = strings.ToLower(strings.TrimSuffix(prefix, "/"))
	regex := regexp.MustCompile(fmt.Sprintf(`^%s-(.+)$`, prefix))

	for _, label := range labels {
		label := strings.ToLower(label)

		if match := regex.FindStringSubmatch(label); match != nil {
			return strings.ToLower(match[1])
		}
	}

	return ""
}
