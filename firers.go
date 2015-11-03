package main

func NewFirer(f float64) NeuronFirer {
	//return SimpleNeuronFirer{}
	return PotentialFirer{pcalc: NewPotentialCalculator(f)}
}

type SimpleNeuronFirer struct {
}

func (_ SimpleNeuronFirer) Fire(neuron *Neuron) {
	neuron.firing = true
}
func (_ SimpleNeuronFirer) UnFire(neuron *Neuron) {
	neuron.firing = false
}

type PotentialFirer struct {
	pcalc PotentialCalculator
}

func (p PotentialFirer) Fire(neuron *Neuron) {
	neuron.firing = true
	neuron.potential = p.pcalc.Calculate(neuron.firing, neuron.potential)
}
func (p PotentialFirer) UnFire(neuron *Neuron) {
	neuron.firing = false
	neuron.potential = p.pcalc.Calculate(neuron.firing, neuron.potential)
}
