package mongodbr

import "go.mongodb.org/mongo-driver/v2/mongo/options"

type copiedOption[T any] struct {
	value *T
}

func (o *copiedOption[T]) List() []func(*T) error {
	if o.value == nil {
		return nil
	}
	return []func(*T) error{
		func(target *T) error {
			*target = *o.value
			return nil
		},
	}
}

func asOptionLister[T any](value *T) options.Lister[T] {
	return &copiedOption[T]{value: value}
}

func asOptionListers[T any](values []*T) []options.Lister[T] {
	if len(values) == 0 {
		return nil
	}
	result := make([]options.Lister[T], 0, len(values))
	for _, value := range values {
		result = append(result, asOptionLister(value))
	}
	return result
}

func ptr[T any](value T) *T {
	return &value
}
