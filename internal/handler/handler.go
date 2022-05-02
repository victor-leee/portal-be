package handler

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/victor-leee/portal-be/internal/image"
	"github.com/victor-leee/portal-be/internal/processor"
	"github.com/victor-leee/portal-be/internal/repo"
	"github.com/victor-leee/portal-be/internal/response_error"
	"google.golang.org/protobuf/proto"
	"net/http"
)

type GinHandler struct {
	Processor      processor.RPCServiceProcessor
	RepoProcessor  repo.Processor
	BuildProcessor image.Processor
}

func (h *GinHandler) CreateService(c *gin.Context) (interface{}, error) {
	var req CreateServiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	return nil, h.Processor.Create(&processor.CreateRPCServiceContext{
		Context:          context.Background(),
		Name:             req.Name,
		HierarchyInfo:    req.Hierarchy,
		ParentID:         req.ParentID,
		IsService:        req.IsService,
		GitRepoURL:       req.GitRepoURL,
		BuildFileRelPath: req.BuildFileRelPath,
	})
}

func (h *GinHandler) QueryByParentID(c *gin.Context) (interface{}, error) {
	var req QueryByParentIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	return h.Processor.QueryByParentID(context.Background(), req.ParentID)
}

func (h *GinHandler) ListBranches(c *gin.Context) (interface{}, error) {
	var req ListBranchesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	s, err := h.Processor.QueryByID(context.Background(), req.ID)
	if err != nil {
		return nil, err
	}
	if len(s.GitRepo) == 0 {
		return nil, &response_error.PortalError{
			InternalError: errors.New("specify git repo url first"),
			ResponseCode:  proto.Int64(http.StatusBadRequest),
		}
	}

	return h.RepoProcessor.ListAllBranches(s.GitRepo)
}

func (h *GinHandler) RunPipeLine(c *gin.Context) (interface{}, error) {
	var req RunPipelineRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	s, err := h.Processor.QueryByID(context.Background(), req.ID)
	if err != nil {
		return nil, err
	}

	// ----- clone repo
	baseDir, err := h.RepoProcessor.Clone(s.GitRepo, req.RemoteBranch)
	if err != nil {
		return nil, err
	}
	// ----- send to docker daemon to build image
	return h.BuildProcessor.BuildAndPush(baseDir, s.BuildFileRelPath)
}
