package main

import (
	"math/rand"
	"time"
)

func NewCalculater(f float64) NeuronCalculater {
	//return GolNeuronCalculater{mapper: NewCartNeuronMapper(-1, 1, -1, 1)}
	//return CountNeuronCalculater{mapper: NewCartNeuronMapper(-2, 2, -2, 2)}
	//return HecticCalculater{mapper: NewCartNeuronMapper(-2, 2, -2, 2), p: 0.00, r: rand.New(rand.NewSource(time.Now().UnixNano()))}
	return PotentialCalculater{
		r:      rand.New(rand.NewSource(time.Now().UnixNano())),
		f:      f,
		direct: NewCartNeuronMapper(-1, -1, -1, -1),
		near:   NewCartNeuronMapper(-2, -1, -2, -1)}
}

type PotentialCalculater struct {
	r      *rand.Rand
	f      float64
	direct NeuronMapper
	near   NeuronMapper
}

func (calc PotentialCalculater) Calculate(neuron Neuron) bool {
	if !neuron.alive {
		return false
	}

	var chaos float64 = 0.00
	var alpha float64 = 0.5
	var beta float64 = 1 - alpha
	var cAlpha float64 = (1 - chaos) * alpha
	var cBeta float64 = (1 - chaos) * beta

	if calc.r.Float64() < chaos {
		return true
	}

	// fire if direct neighbor is firing with some probability
	for _, direct := range calc.direct.Ins(neuron) {
		if direct.firing && calc.r.Float64() < cAlpha {
			return true
		}
	}

	var pot float64 = 0
	var neighbors = calc.near.Ins(neuron)
	for _, neighbor := range neighbors {
		pot += neighbor.potential
	}
	pot /= float64(len(neighbors))
	temp := 1.0
	return calc.r.Float64() < cBeta*(pot*2)*calc.f*temp
}

type GolNeuronCalculater struct {
	mapper NeuronMapper
}

func (gol GolNeuronCalculater) Calculate(neuron Neuron) bool {
	count := 0
	for _, neighbor := range gol.mapper.Ins(neuron) {
		if neighbor.firing {
			count += 1
		}
	}
	return (neuron.firing && count >= 2 && count <= 3) ||
		(!neuron.firing && count == 3)
}

type CountNeuronCalculater struct {
	mapper NeuronMapper
}

func (calc CountNeuronCalculater) Calculate(neuron Neuron) bool {
	count := 0
	for _, neighbor := range calc.mapper.Ins(neuron) {
		if neighbor.firing {
			count += 1
		}
	}
	return (neuron.firing && count >= 2 && count <= 3) ||
		(!neuron.firing && count == 3)
}

type HecticCalculater struct {
	mapper NeuronMapper
	p      float64
	r      *rand.Rand
}

func (calc HecticCalculater) Calculate(neuron Neuron) bool {
	count := 0
	for _, neighbor := range calc.mapper.Ins(neuron) {
		if neighbor.firing {
			count += 1
		}
	}
	return (calc.r.Float64() < calc.p) ||
		(neuron.firing && count >= 2 && count <= 3) ||
		(!neuron.firing && count == 3)
}
