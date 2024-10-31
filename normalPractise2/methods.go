package main

type Vertex struct {
	X, Y int
}

func (v Vertex) Multiply() int {
	return v.X * v.Y
}

var v = Vertex{2, 5}