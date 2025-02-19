package repository

import (
	"errors"
	"piefiredie/internal/errorcollection"
)

func CreateBeefStock() BeefStock {
	return BeefStock{
		Stock: make(map[string]int32, 0),
	}
}

func (r *BeefStock) AddToStock(input string, count ...int32) (int32, error) {
	if len(count) > 1 {
		return int32(0), errors.Join(errorcollection.ErrRepository, ErrAddToStock, errors.New("Multiple value input"))
	}
	if len(count) != 0 && count[0] < 0 {
		return int32(0), errors.Join(errorcollection.ErrRepository, ErrAddToStock, errors.New("Invalid value input"))
	}

	addcount := int32(1)
	if len(count) == 1 {
		addcount = count[0]
	}

	v, ok := r.Stock[input]
	if !ok {
		r.Stock[input] = addcount
	} else {
		r.Stock[input] = v + addcount
	}

	return v + addcount, nil
}

func (r *BeefStock) RemoveFromStock(input string, count ...int32) (int32, error) {
	if len(count) > 1 {
		return int32(0), errors.Join(errorcollection.ErrRepository, ErrRemoveFromStock, errors.New("Multiple value input"))
	}
	if len(count) != 0 && count[0] < 0 {
		return int32(0), errors.Join(errorcollection.ErrRepository, ErrRemoveFromStock, errors.New("Invalid value input"))
	}

	removecount := int32(1)
	if len(count) == 1 {
		removecount = count[0]
	}

	v, ok := r.Stock[input]
	if !ok {
		return int32(0), errors.Join(errorcollection.ErrRepository, ErrRemoveFromStock, errors.New("Stock not found"))
	}
	if v < removecount {
		return int32(0), errors.Join(errorcollection.ErrRepository, ErrRemoveFromStock, errors.New("Stock not enough"))
	}

	r.Stock[input] = v - removecount

	return v - removecount, nil
}

func (r *BeefStock) GetStock(opts ...string) (map[string]int32, error) {
	if len(opts) > 0 {
		ret := make(map[string]int32, 0)
		for _, v := range opts {
			stock, ok := r.Stock[v]
			if !ok {
				return nil, errors.Join(errorcollection.ErrRepository, ErrGetStock, errors.New("Unknown beef"))
			} else {
				ret[v] = stock
			}
		}
		return ret, nil
	}

	return r.Stock, nil
}
