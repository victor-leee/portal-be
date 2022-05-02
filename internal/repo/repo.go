package repo

type Processor interface {
	ListAllBranches(repoURL string) ([]string, error)
	Clone(repoURL, branch string) (string, error)
}
