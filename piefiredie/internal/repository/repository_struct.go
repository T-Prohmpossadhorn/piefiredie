package repository

import "errors"

type BeefStock struct {
	Stock map[string]int32
}

var (
	ErrAddToStock      = errors.New("AddToStock error")
	ErrRemoveFromStock = errors.New("RemoveFromStock error")
	ErrGetStock        = errors.New("GetStock error")
)
