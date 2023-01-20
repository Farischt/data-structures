package heap

type Item[T number] struct {
	Value       T
	Information interface{}
}

func NewItem[T number](v T, info interface{}) *Item[T] {
	return &Item[T]{
		Value:       v,
		Information: info,
	}
}
