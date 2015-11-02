package main

type Firer interface {
	Fire(node Node)
	UnFire(node Node)
}

type FireCalculater interface {
	Calculate(neuron Neuron) (fire bool)
}

type NetCreater interface {
	Create(n int, f float64) (net Net)
}

type NetPrinter interface {
	Print(net Net)
}
