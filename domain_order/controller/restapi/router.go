package restapi

import (
	"github.com/gin-gonic/gin"

	"demo4/domain_order/usecase/createorder"
	"demo4/shared/infrastructure/config"
	"demo4/shared/infrastructure/logger"
)

type Controller struct {
	Router            gin.IRouter
	Config            *config.Config
	Log               logger.Logger
	CreateOrderInport createorder.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	r.Router.POST("/createorder", r.authorized(), r.createOrderHandler(r.CreateOrderInport))
}
