package application

import (
	"demo4/domain_order/controller/restapi"
	"demo4/domain_order/gateway/prod"
	"demo4/domain_order/usecase/createorder"
	"demo4/shared/driver"
	"demo4/shared/infrastructure/config"
	"demo4/shared/infrastructure/logger"
	"demo4/shared/infrastructure/server"
	"demo4/shared/infrastructure/util"
)

type demo4 struct {
	httpHandler *server.GinHTTPHandler
	controller  driver.Controller
}

func (c demo4) RunApplication() {
	c.controller.RegisterRouter()
	c.httpHandler.RunApplication()
}

func NewDemo4() func() driver.RegistryContract {
	return func() driver.RegistryContract {

		cfg := config.ReadConfig()

		appID := util.GenerateID(4)

		appData := driver.NewApplicationData("demo4", appID)

		log := logger.NewSimpleJSONLogger(appData)

		httpHandler := server.NewGinHTTPHandlerDefault(log, appData, cfg)

		datasource := prod.NewGateway(log, appData, cfg)

		return &demo4{
			httpHandler: &httpHandler,
			controller: &restapi.Controller{
				Log:               log,
				Config:            cfg,
				Router:            httpHandler.Router,
				CreateOrderInport: createorder.NewUsecase(datasource),
			},
		}

	}
}
