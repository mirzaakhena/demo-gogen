package prod

import (
	"context"
	"demo4/domain_order/model/entity"
	"demo4/shared/driver"
	"demo4/shared/infrastructure/config"
	"demo4/shared/infrastructure/logger"
)

type gateway struct {
	log     logger.Logger
	appData driver.ApplicationData
	config  *config.Config
}

// NewGateway ...
func NewGateway(log logger.Logger, appData driver.ApplicationData, config *config.Config) *gateway {

	return &gateway{
		log:     log,
		appData: appData,
		config:  config,
	}
}

func (r *gateway) SaveOrder(ctx context.Context, obj *entity.Order) error {
	r.log.Info(ctx, "called")

	return nil
}
