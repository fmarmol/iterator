package iterator

type List[T any] []T

func NewList[T any](elems ...T) List[T] {
	r := make(List[T], len(elems))
	for index, elem := range elems {
		r[index] = elem
	}
	return r
}

func (l *List[T]) Iter() <-chan T {
	ret := make(chan T)
	go func() {
		for _, elem := range *l {
			ret <- elem

		}
		close(ret)
	}()
	return ret
}
