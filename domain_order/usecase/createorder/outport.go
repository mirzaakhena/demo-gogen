package createorder

import "demo4/domain_order/model/repository"

// Outport of usecase
type Outport interface {
	repository.SaveOrderRepo
}
