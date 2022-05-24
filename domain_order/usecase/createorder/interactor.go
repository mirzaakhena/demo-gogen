package createorder

import (
	"context"
	"demo4/domain_order/model/entity"
)

//go:generate mockery --name Outport -output mocks/

type createOrderInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &createOrderInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *createOrderInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	res := &InportResponse{}

	// code your usecase definition here ...

	orderObj, err := entity.NewOrder(entity.OrderRequest{})
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveOrder(ctx, orderObj)
	if err != nil {
		return nil, err
	}

	//!

	return res, nil
}
