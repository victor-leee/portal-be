package handler

type CreateServiceRequest struct {
	Name      string   `json:"name"`
	Hierarchy []string `json:"hierarchy"`
	ParentID  uint64   `json:"parentID"`
	IsService bool     `json:"isService"`
}

type QueryByParentIDRequest struct {
	ParentID uint64 `json:"parentID"`
}
