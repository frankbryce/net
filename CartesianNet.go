package main

type GridNet struct {
	Net
	Neurons [][]Neuron
}

// will create GridNets
type CartesianNet struct {
}

func (_ CartesianNet) Create(n int, f float64) (net Net) {
	return
}

func (_ CartesianNet) Print(net Net) {
	return
}

func (_ CartesianNet) Fire(node Node) {
	return
}

func (_ CartesianNet) UnFire(node Node) {
	return
}

func (_ CartesianNet) Calculate(neuron Neuron) (fire bool) {
	return
}
