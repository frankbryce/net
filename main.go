package main

import (
	"os"
	"os/exec"
	"sync"
	"time"
)

var size int = 30
var factor float64 = 0.1
var steps int = 10000
var delay = 200 * time.Millisecond

func main() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	simulate(
		NewCartMaker(factor),
		CartCloner{},
		NewCalculater(factor),
		NewFirer(factor),
		NewPrinter())
}

func calc(in Neuron, out *Neuron, nc NeuronCalculater, nf NeuronFirer, wg *sync.WaitGroup) {
	defer wg.Done()
	if nc.Calculate(in) {
		nf.Fire(out)
	} else {
		nf.UnFire(out)
	}
}

func step(in Net, out Net, nc NeuronCalculater, nf NeuronFirer) {
	var wg sync.WaitGroup
	for i := range in.neurons {
		wg.Add(len(in.neurons[i]))
	}
	for i := range in.neurons {
		for j, neuron := range in.neurons[i] {
			go calc(neuron, &(out.neurons[i][j]), nc, nf, &wg)
		}
	}
	wg.Wait()
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
		time.Sleep(delay)
		np.Print(nets[curr])
	}
}
