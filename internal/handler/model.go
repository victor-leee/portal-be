package handler

type CreateServiceRequest struct {
	Name      string   `json:"name" binding:"required"`
	Hierarchy []string `json:"hierarchy" binding:"required"`
	ParentID  uint64   `json:"parentID" binding:"required"`
	IsService bool     `json:"isService" binding:"required"`
}

type QueryByParentIDRequest struct {
	ParentID uint64 `json:"parentID" binding:"required"`
}
