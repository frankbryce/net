package main

import (
	"fmt"
	"math"
)

func NewPrinter() NetPrinter {
	return AliveCartPrinter{}
	//return PotentialPrinter{}
}

type FullCartPrinter struct {
}

func (_ FullCartPrinter) Print(net Net) {
	fmt.Printf("\033[0;0H")
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

type AliveCartPrinter struct {
}

func (_ AliveCartPrinter) Print(net Net) {
	fmt.Printf("\033[0;0H")
	for i := range net.neurons {
		for _, neuron := range net.neurons[i] {
			if neuron.firing {
				fmt.Printf("*")
			} else if neuron.alive {
				fmt.Printf(".")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\r")
}

type PotentialPrinter struct {
}

func (_ PotentialPrinter) Print(net Net) {
	fmt.Printf("\033[0;0H")
	for i := range net.neurons {
		for _, neuron := range net.neurons[i] {
			if !neuron.alive {
				fmt.Printf(" ")
			} else {
				digitPotential := Min(int64(math.Floor(neuron.potential*100)), 99)
				if digitPotential != 0 {
					fmt.Printf("%2d ", digitPotential)
				} else {
					fmt.Printf("-- ")
				}
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\r")
}
