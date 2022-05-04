package handler

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	config_backend "github.com/victor-leee/portal-be/gen/github.com/victor-leee/config-backend"
	"github.com/victor-leee/portal-be/internal/cluster"
	"github.com/victor-leee/portal-be/internal/image"
	"github.com/victor-leee/portal-be/internal/processor"
	"github.com/victor-leee/portal-be/internal/processor/pipeline"
	"github.com/victor-leee/portal-be/internal/repo"
	"github.com/victor-leee/portal-be/internal/response_error"
	"github.com/victor-leee/portal-be/internal/rpc"
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
		Type:             req.Type,
		CustomPort:       req.CustomPort,
		PrefixMapping:    req.PrefixMapping,
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

	var progressErr error
	var baseDir string
	var tag *string

	return pipeline.New().
		Clone(func() error {
			baseDir, progressErr = h.RepoProcessor.Clone(s.GitRepo, req.RemoteBranch)
			return progressErr
		}).
		Build(func() error {
			tag, progressErr = h.BuildProcessor.BuildAndPush(baseDir, s.BuildFileRelPath)
			if progressErr != nil {
				return progressErr
			}
			if tag == nil {
				return errors.New("empty tag")
			}

			return nil
		}).
		Push(func() error {
			return nil
		}).
		Deploy(func() error {
			return cluster.GetManager(cluster.K8S).ApplyServiceDeployment(context.Background(), &cluster.DeploymentConfig{
				Service:  s,
				Replicas: req.Replicas,
				ImageTag: *tag,
			})
		}).Do(context.Background())

}

func (h *GinHandler) QueryPipelineStatusByID(c *gin.Context) (interface{}, error) {
	var req QueryPipelineStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	return pipeline.Query(req.PipelineRunID), nil
}

func (h *GinHandler) PutConfig(c *gin.Context) (interface{}, error) {
	var req PutConfigRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	return rpc.ConfigBackendCli.PutConfig(context.Background(), &config_backend.PutConfigRequest{
		ServiceId:  req.ServiceID,
		ServiceKey: req.ServiceKey,
		Key:        req.Key,
		Value:      req.Value,
	})
}

func (h *GinHandler) GetConfig(c *gin.Context) (interface{}, error) {
	var req GetConfigRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	return rpc.ConfigBackendCli.GetConfig(context.Background(), &config_backend.GetConfigRequest{
		ServiceId:  req.ServiceID,
		ServiceKey: req.ServiceKey,
		Key:        req.Key,
	})
}

func (h *GinHandler) GetConfigKeys(c *gin.Context) (interface{}, error) {
	var req GetConfigRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	return rpc.ConfigBackendCli.GetAllKeys(context.Background(), &config_backend.GetAllKeysRequest{
		ServiceId:  req.ServiceID,
		ServiceKey: req.ServiceKey,
	})
}
