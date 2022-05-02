package repo

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/sirupsen/logrus"
	"github.com/victor-leee/portal-be/internal/config"
	"net/http"
	"os"
	"strings"
	"time"
)

type Branch struct {
	Name string `json:"name"`
}

type GithubProcessor struct {
}

func (d *GithubProcessor) Clone(repoURL, branch string) (string, error) {
	path := fmt.Sprintf("/%s/%s/%s/%d", os.TempDir(), base64.URLEncoding.EncodeToString([]byte(repoURL)),
		branch, time.Now().UnixMilli())
	logrus.Infof("local path:%s", path)
	_, err := git.PlainClone(path, false,
		&git.CloneOptions{
			URL:           repoURL,
			ReferenceName: plumbing.NewBranchReferenceName(branch),
			SingleBranch:  true,
		})
	return path, err
}

func (d *GithubProcessor) ListAllBranches(repoURL string) ([]string, error) {
	user, repo := d.refineUserAndRepo(repoURL)
	var branches []*Branch
	resp, err := http.Get(fmt.Sprintf(config.APIRetrieveBranches, user, repo))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err = json.NewDecoder(resp.Body).Decode(&branches); err != nil {
		return nil, err
	}

	branchNames := make([]string, 0, len(branches))
	for _, b := range branches {
		branchNames = append(branchNames, b.Name)
	}
	return branchNames, nil
}

func (d *GithubProcessor) refineUserAndRepo(repoURL string) (username, repoName string) {
	repoURL = strings.TrimLeft(repoURL, "https://")
	repoURL = strings.TrimLeft(repoURL, "http://")
	repoURL = strings.TrimRight(repoURL, ".git")
	// repoURL == github.com/{username}/{repoName}
	components := strings.Split(repoURL, "/")

	return components[1], components[2]
}
