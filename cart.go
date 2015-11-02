// Cartesian x/y grid view on the neural network
package main

import (
	"fmt"
	"math/rand"
)

type CartMaker struct {
}

func (cm CartMaker) Make(n int, f float64) Net {
	net := Net{neurons: make([][]Neuron, n, n)}
	for i := 0; i < n; i++ {
		net.neurons[i] = make([]Neuron, n, n)
		for j := 0; j < n; j++ {
			net.neurons[i][j] = Neuron{
				firing: false,
				net:    net,
				x:      i,
				y:      j,
			}
			if rand.Float64() < f {
				net.neurons[i][j].firing = true
			}
		}
	}
	return net
}

type CartCloner struct {
}

func (c CartCloner) Clone(in Net) Net {
	n := len(in.neurons)
	out := Net{neurons: make([][]Neuron, n, n)}
	for i := 0; i < n; i++ {
		out.neurons[i] = make([]Neuron, n, n)
		for j := 0; j < n; j++ {
			out.neurons[i][j] = Neuron{
				potential: in.neurons[i][j].potential,
				firing:    in.neurons[i][j].firing,
				net:       out,
				x:         i,
				y:         j,
			}
		}
	}
	return out
}

type CartNeuronMapper struct {
}

func (c CartNeuronMapper) Ins(neuron Neuron) []Neuron {
	return c.Outs(neuron)
}
func (_ CartNeuronMapper) Outs(neuron Neuron) []Neuron {
	neurons := make([]Neuron, 8, 8)
	x, y := neuron.x, neuron.y
	xm, xp, ym, yp :=
		(x+len(neuron.net.neurons)-1)%len(neuron.net.neurons),
		(x+1)%len(neuron.net.neurons),
		(y+len(neuron.net.neurons)-1)%len(neuron.net.neurons[x]),
		(y+1)%len(neuron.net.neurons[x])
	neurons[0] = neuron.net.neurons[xm][y]
	neurons[1] = neuron.net.neurons[x][yp]
	neurons[2] = neuron.net.neurons[xp][y]
	neurons[3] = neuron.net.neurons[x][ym]
	neurons[4] = neuron.net.neurons[xm][ym]
	neurons[5] = neuron.net.neurons[xm][yp]
	neurons[6] = neuron.net.neurons[xp][ym]
	neurons[7] = neuron.net.neurons[xp][yp]
	return neurons
}

type SimpleNeuronFirer struct {
}

func (_ SimpleNeuronFirer) Fire(neuron *Neuron) {
	neuron.firing = true
}
func (_ SimpleNeuronFirer) UnFire(neuron *Neuron) {
	neuron.firing = false
}

type GolNeuronCalculater struct {
}

func (_ GolNeuronCalculater) Calculate(neuron Neuron) bool {
	mapper := CartNeuronMapper{}
	count := 0
	for _, neighbor := range mapper.Ins(neuron) {
		if neighbor.firing {
			count += 1
		}
	}
	return (neuron.firing && count >= 2 && count <= 3) ||
		(!neuron.firing && count == 3)
}

type FullCartPrinter struct {
}

func (cp FullCartPrinter) Print(net Net) {
	fmt.Printf("\033[0;0H")
	fmt.Printf("-------------------\n")
	for i := range net.neurons {
		for _, neuron := range net.neurons[i] {
			if neuron.firing {
				fmt.Printf("X")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\r")
}
