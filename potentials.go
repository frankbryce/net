package main

import "math"

func NewPotentialCalculator(f float64) PotentialCalculator {
	//return SigmoidCalculatorAlpha{alpha: 0.15, sgm: NewSigmoid(), nzr: NewNormalizer(f)}
	return SigmoidCalculator{f: f, delta: 0.05 * f, sgm: NewSigmoid()}
}

type Sigmoid interface {
	sgm(in float64) (out float64)
}

type Normalizer interface {
	norm(p float64) float64
}

type SigmoidCalculatorAlpha struct {
	alpha float64
	sgm   Sigmoid
	nzr   Normalizer
}

type SigmoidCalculator struct {
	f     float64
	delta float64
	sgm   Sigmoid
}

func (s SigmoidCalculator) Calculate(fired bool, pot float64) (newPot float64) {
	if fired {
		newPot = s.sgm.sgm(pot + math.Max(1.0, (1.0-s.f)/s.f)*s.delta)
	} else {
		newPot = s.sgm.sgm(pot - math.Max(1.0, s.f/(1.0-s.f))*s.delta)
	}
	return
}

func NewSigmoid() Sigmoid {
	//return Sigmoid1{}
	return Linear{}
}

type Linear struct{}

func (l Linear) sgm(in float64) float64 {
	return math.Min(math.Max(in, 0.0), 1.0)
}

func NewNormalizer(f float64) Normalizer {
	//return Normalizer1{exp: math.Log(0.5) / math.Log(f)}
	return Normalizer2{}
}

type Normalizer1 struct {
	exp float64
}

func (nzr Normalizer1) norm(in float64) float64 {
	return math.Pow(in, nzr.exp)
}

type Normalizer2 struct {
}

func (nzr Normalizer2) norm(in float64) float64 {
	return in
}

type Sigmoid1 struct {
}

func (s Sigmoid1) sgm(in float64) (out float64) {
	t := 2*in - 1
	return math.Min(math.Max(t/(1+math.Abs(t))+0.5, 0), 1)
}

func (s SigmoidCalculatorAlpha) Calculate(fired bool, pot float64) (newPot float64) {
	if fired {
		newPot = s.nzr.norm(s.sgm.sgm(s.alpha + pot*(1-s.alpha)))
	} else {
		newPot = s.nzr.norm(s.sgm.sgm(pot * (1 - s.alpha)))
	}
	return
}
