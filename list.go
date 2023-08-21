package collection

type List[T any] interface {
	Add(T)

	AddAll(...T)

	Insert(int, T)
}
