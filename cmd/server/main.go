package main

import (
	errors2 "errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/victor-leee/portal-be/internal/config"
	"github.com/victor-leee/portal-be/internal/errors"
	"github.com/victor-leee/portal-be/internal/handler"
	"github.com/victor-leee/portal-be/internal/model"
	"github.com/victor-leee/portal-be/internal/processor"
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
	logrus.Info("starting service")
	r := gin.Default()
	h := handler.GinHandler{
		Processor: &processor.DefaultRPCServiceProcessor{
			ServiceDao: &model.RPCServiceDaoImpl{},
		},
	}
	r.POST("/create-service", wrapperHandler(h.CreateService))
	r.POST("/query-by-parent-id", wrapperHandler(h.QueryByParentID))
	logrus.Fatal(r.Run(":80"))
}

func wrapperHandler(f CustomHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		message, err := f(ctx)
		portalErr := &errors.PortalError{}
		if err == nil {
			ctx.JSONP(http.StatusOK, &ResponseMessage{
				Code:    http.StatusOK,
				Message: message,
			})
			return
		}

		if errors2.As(err, &portalErr) {
			ctx.JSONP(int(portalErr.Code()), err.Error())
			return
		}
		ctx.JSONP(http.StatusInternalServerError, err.Error())
	}
}
