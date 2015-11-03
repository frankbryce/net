package main

type CartCloner struct {
}

func (c CartCloner) Clone(in Net) Net {
	n := len(in.neurons)
	out := Net{neurons: make([][]Neuron, n, n)}
	for i := 0; i < n; i++ {
		out.neurons[i] = make([]Neuron, n, n)
		for j := 0; j < n; j++ {
			out.neurons[i][j] = Neuron{
				alive:     in.neurons[i][j].alive,
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
