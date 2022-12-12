package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/okieraised/xelnaga/internal"
	"go.uber.org/zap"
)

type instancesAPIRoute struct {
	engine *gin.Engine
	logger *zap.Logger
}

func NewInstancesRoute(engine *gin.Engine, logger *zap.Logger) {
	app := &instancesAPIRoute{
		engine: engine,
		logger: logger,
	}
	app.init()
}

func (r *instancesAPIRoute) init() {
	instancesRoute := r.engine.Group("/instances")

	instancesRoute.GET("", r.listInstances)
	instancesRoute.GET(fmt.Sprintf("/%s", internal.SOPInstanceUID), r.getInstance)
	instancesRoute.DELETE(fmt.Sprintf("/%s", internal.SOPInstanceUID), r.getInstance)

}

func (r *instancesAPIRoute) listInstances(ctx *gin.Context) {
}

func (r *instancesAPIRoute) deleteInstance(ctx *gin.Context) {
}

func (r *instancesAPIRoute) getInstance(ctx *gin.Context) {
}
