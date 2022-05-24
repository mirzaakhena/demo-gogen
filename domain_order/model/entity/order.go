package entity

type Order struct {
	ID string `` //
}

type OrderRequest struct {
}

func NewOrder(req OrderRequest) (*Order, error) {

	var obj Order

	// assign value here

	err := obj.Validate()
	if err != nil {
		return nil, err
	}

	return &obj, nil
}

func (r *Order) Validate() error {
	return nil
}
