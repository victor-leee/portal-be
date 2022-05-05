package handler

type CreateServiceRequest struct {
	Name             string   `json:"name"`
	Hierarchy        []string `json:"hierarchy"`
	ParentID         uint64   `json:"parentID"`
	IsService        bool     `json:"isService"`
	GitRepoURL       string   `json:"gitRepoURL"`
	BuildFileRelPath string   `json:"buildFileRelPath"`
	Type             string   `json:"type"`
	CustomPort       uint16   `json:"customPort"`
	PrefixMapping    string   `json:"prefixMapping"`
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

type PutConfigRequest struct {
	ServiceID  string `json:"serviceID"`
	ServiceKey string `json:"serviceKey"`
	Key        string `json:"key"`
	Value      string `json:"value"`
}

type GetConfigRequest struct {
	ServiceID  string `json:"serviceID"`
	ServiceKey string `json:"serviceKey"`
	Key        string `json:"key"`
}

type SCRPCBaseResponse struct {
	ErrCode int    `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
}

type GetConfigResponse struct {
	BaseResponse *SCRPCBaseResponse `json:"baseResponse"`
	KeyExist     bool               `json:"keyExist"`
	Value        string             `json:"value"`
}

type PutConfigResponse struct {
	BaseResponse *SCRPCBaseResponse `json:"baseResponse"`
}

type GetConfigKeysResponse struct {
	BaseResponse *SCRPCBaseResponse `json:"baseResponse"`
	Keys         []string           `json:"keys"`
}
