package heap

import (
	"github.com/farischt/ds/utils"
)

type Item[T utils.Number] struct {
	Value       T
	Information interface{}
}

func NewItem[T utils.Number](v T, info interface{}) *Item[T] {
	return &Item[T]{
		Value:       v,
		Information: info,
	}
}
