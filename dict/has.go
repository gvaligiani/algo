package dict

// //////////////////////////////////////////////////
// has

func Has[K comparable, V comparable](items map[K]V, value V) bool {
	for _, item := range items {
		if item == value {
			return true
		}
	}
	return false
}