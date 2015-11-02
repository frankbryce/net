package main

type Neuron struct {
	potential float64
	firing    bool
	net       Net
	x, y      int
}

type Net struct {
	neurons [][]Neuron
}

type NetMaker interface {
	Make(n int, f float64) Net
}

type NetCloner interface {
	Clone(net Net) Net
}

type NeuronMapper interface {
	Ins(neuron Neuron) []Neuron
	Outs(neuron Neuron) []Neuron
}

type NeuronFirer interface {
	Fire(neuron *Neuron)
	UnFire(neuron *Neuron)
}

type NeuronCalculater interface {
	Calculate(neuron Neuron) bool
}

type NetPrinter interface {
	Print(net Net)
}
