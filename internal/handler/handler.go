package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/victor-leee/portal-be/internal/processor"
)

type GinHandler struct {
	Processor processor.RPCServiceProcessor
}

func (h *GinHandler) CreateService(c *gin.Context) (interface{}, error) {
	var req CreateServiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	return nil, h.Processor.Create(context.Background(), req.Name, req.Hierarchy, req.ParentID, req.IsService)
}

func (h *GinHandler) QueryByParentID(c *gin.Context) (interface{}, error) {
	var req QueryByParentIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	return h.Processor.Query(context.Background(), req.ParentID)
}
