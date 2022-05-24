package repository

import (
	"context"
	"demo4/domain_order/model/entity"
)

type SaveOrderRepo interface {
	SaveOrder(ctx context.Context, obj *entity.Order) error
}
