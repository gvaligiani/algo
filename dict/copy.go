package dict

func Copy[K comparable, T any, M ~map[K]T](items M) M {
	return CopyIf(items, True[K, T]())
}