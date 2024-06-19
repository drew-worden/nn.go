package main

import (
	"fmt"
)

func main() {

	inputs := []float64{1, 2, 3, 2.5}

	n1 := Neuron{Weights: []float64{0.2, 0.8, -0.5, 1}, Bias: 2}
	n2 := Neuron{Weights: []float64{0.5, -0.91, 0.26, -0.5}, Bias: 3}
	n3 := Neuron{Weights: []float64{-0.26, -0.27, 0.17, 0.87}, Bias: 0.5}

	neurons := []Neuron{n1, n2, n3}
	var outputs []float64

	for i := 0; i < len(neurons); i++ {
		neurons := neurons[i]
		output := DotProduct(inputs, neurons.Weights) + neurons.Bias
		outputs = append(outputs, output)
	}
	fmt.Println(outputs)
}
