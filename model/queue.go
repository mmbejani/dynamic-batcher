package model

type Node[T any] struct {
	id   int16
	data []T
}

type Queue[T any] struct {
	root   *Node[T]
	tail   *Node[T]
	length int16
}
