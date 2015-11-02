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

type Net struct {
	Nodes []Node
}
