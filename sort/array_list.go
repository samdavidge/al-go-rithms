package sort

import "errors"

const BaseCapacity = 20

type ArrayList[T interface{}] struct {
	values []T
	length int
}

func (arrayList *ArrayList[T]) Get(index int) (T, error) {
	var value T
	if index > arrayList.length-1 {
		return value, errors.New("The array list is empty")
	}

	value = arrayList.values[index]

	return value, nil
}

func (arrayList *ArrayList[T]) Push(value T) {
	arrayList.ensureCapacity()
	arrayList.values[arrayList.length] = value
	arrayList.length++
}

func (arrayList *ArrayList[T]) Pop() (T, error) {
	var value T
	if arrayList.length == 0 {
		return value, errors.New("The array list is empty")
	}

	value = arrayList.values[arrayList.length-1]

	return value, nil
}

// I could just use append(), but we're managing/reserving our
// own capacity as per the lesson
func (arrayList *ArrayList[T]) ensureCapacity() {
	if arrayList.length == len(arrayList.values) {
		arrayList.values = append(arrayList.values, make([]T, BaseCapacity)...)
	}
}
