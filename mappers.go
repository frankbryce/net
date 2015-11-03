package main

type CartNeuronMapper struct {
	xm, xp, ym, yp int
}

func NewCartNeuronMapper(xm, xp, ym, yp int) CartNeuronMapper {
	return CartNeuronMapper{xm: xm, xp: xp, ym: ym, yp: yp}
}

func (c CartNeuronMapper) Ins(neuron Neuron) []Neuron {
	return c.Outs(neuron)
}
func (mpr CartNeuronMapper) Outs(neuron Neuron) []Neuron {
	neurons := make([]Neuron, 0, 8)
	x, y := neuron.x, neuron.y
	xmod := len(neuron.net.neurons)
	xm, xp, ym, yp :=
		(x+xmod+mpr.xm)%xmod,
		(x+mpr.xp)%xmod,
		(y+len(neuron.net.neurons[x])+mpr.ym)%len(neuron.net.neurons[x]),
		(y+mpr.yp)%len(neuron.net.neurons[x])
	for i := xm; i != (xp+1)%xmod; i = (i + 1) % xmod {
		ymod := len(neuron.net.neurons[i])
		for j := ym; j != (yp+1)%ymod; j = (j + 1) % ymod {
			if i == x && j == y {
				continue
			}
			if neuron.net.neurons[i][j].alive {
				neurons = append(neurons, neuron.net.neurons[i][j])
			}
		}
	}
	return neurons
}
