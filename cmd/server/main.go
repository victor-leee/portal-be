package main

import (
	errors2 "errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/victor-leee/portal-be/internal/config"
	"github.com/victor-leee/portal-be/internal/handler"
	"github.com/victor-leee/portal-be/internal/image"
	"github.com/victor-leee/portal-be/internal/model"
	"github.com/victor-leee/portal-be/internal/processor"
	"github.com/victor-leee/portal-be/internal/repo"
	"github.com/victor-leee/portal-be/internal/response_error"
	"net/http"
)

type CustomHandler func(ctx *gin.Context) (interface{}, error)
type ResponseMessage struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func main() {
	cfg, err := config.Init()
	if err != nil {
		panic(err)
	}
	model.MustInit(cfg)
	image.MustInitBuildCfg(cfg)
	r := gin.Default()
	h := handler.GinHandler{
		Processor: &processor.DefaultRPCServiceProcessor{
			ServiceDao: &model.RPCServiceDaoImpl{},
		},
		RepoProcessor:  &repo.GithubProcessor{},
		BuildProcessor: &image.Docker{},
	}
	r.POST("/create-service", wrapperHandler(h.CreateService))
	r.POST("/query-by-parent-id", wrapperHandler(h.QueryByParentID))
	r.POST("/list-branches", wrapperHandler(h.ListBranches))
	r.POST("/run-pipeline", wrapperHandler(h.RunPipeLine))
	r.POST("/pipeline-stage", wrapperHandler(h.QueryPipelineStatusByID))
	logrus.Fatal(r.Run(":80"))
}

func wrapperHandler(f CustomHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		message, err := f(ctx)
		portalErr := &response_error.PortalError{}
		if err == nil {
			ctx.JSONP(http.StatusOK, &ResponseMessage{
				Code:    http.StatusOK,
				Message: message,
			})
			return
		}

		var code int
		var msg string
		if errors2.As(err, &portalErr) {
			code, msg = int(portalErr.Code()), portalErr.Error()
		} else {
			code, msg = http.StatusInternalServerError, err.Error()
		}
		ctx.JSONP(code, &ResponseMessage{
			Code:    code,
			Message: msg,
		})
	}
}
