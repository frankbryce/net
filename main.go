package main

import (
	"os"
	"os/exec"
	"time"
)

var size int = 20
var factor float64 = 0.25
var steps int = 10000
var delay = 40 * time.Millisecond

func main() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	simulate(
		CartMaker{},
		CartCloner{},
		GolNeuronCalculater{},
		SimpleNeuronFirer{},
		FullCartPrinter{})
}

func step(in Net, out Net, nc NeuronCalculater, nf NeuronFirer) {
	for i := range in.neurons {
		for j, neuron := range in.neurons[i] {
			if nc.Calculate(neuron) {
				nf.Fire(&out.neurons[i][j])
			} else {
				nf.UnFire(&out.neurons[i][j])
			}
		}
	}
	return
}

func simulate(nm NetMaker, ncl NetCloner, nca NeuronCalculater, nf NeuronFirer, np NetPrinter) {
	var nets [2]Net
	nets[0] = nm.Make(size, factor)
	nets[1] = ncl.Clone(nets[0])
	var curr int = 0
	var next int = 1
	np.Print(nets[curr])
	for i := 1; i < steps; i += 1 {
		step(nets[curr], nets[next], nca, nf)
		next, curr = curr, next
		np.Print(nets[curr])
		time.Sleep(delay)
	}
}
