package gox

func GetAttribute[T any](n Node, name string) T {
	for _, a := range n.(node).attributes {
		if a.name == name {
			return a.value.(T)
		}
	}
	return *new(T)
}
