package gox

func CreateAttribute[T any](name string) func(value ...T) Node {
	return func(value ...T) Node {
		return createAttribute[T](name, value...)
	}
}

func CreateElement(name string) func(nodes ...Node) Node {
	return func(nodes ...Node) Node {
		return createElement(name, nodes...)
	}
}

func CreateShared(name string) func(nodes ...Node) Node {
	return func(nodes ...Node) Node {
		return createShared(name, nodeElement, nodes...)
	}
}
