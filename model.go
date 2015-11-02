package main

type Neuron struct {
	Potential float64
	IsFiring  bool
}

type Node struct {
	Neuron
	Id   int
	Ins  []Node
	Outs []Node
}

type Grid struct {
	Neurons [][]Neuron
}

type Graph struct {
	Nodes []Node
}

type Net struct {
	Graph
	Grid
}
