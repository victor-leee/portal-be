package handler

type CreateServiceRequest struct {
	Name             string   `json:"name"`
	Hierarchy        []string `json:"hierarchy"`
	ParentID         uint64   `json:"parentID"`
	IsService        bool     `json:"isService"`
	GitRepoURL       string   `json:"gitRepoURL"`
	BuildFileRelPath string   `json:"buildFileRelPath"`
}

type QueryByParentIDRequest struct {
	ParentID uint64 `json:"parentID"`
}

type ListBranchesRequest struct {
	ID uint64 `json:"id"`
}

type RunPipelineRequest struct {
	ID           uint64 `json:"id"`
	RemoteBranch string `json:"remoteBranch"`
	Replicas     int32  `json:"replicas"`
}

type QueryPipelineStatusRequest struct {
	PipelineRunID string `json:"ID"`
}
