package main

func DotProduct(tensor1 []float64, tensor2 []float64) float64 {
	var output float64
	for i := 0; i < len(tensor1); i++ {
		output += tensor1[i] * tensor2[i]
	}
	return output
}
