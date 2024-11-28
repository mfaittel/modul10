package main

import "fmt"

func main() {
	var N int
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&N)

	weights := make([]float64, N)

	fmt.Printf("Masukkan berat %d anak kelinci:\n", N)
	for i := 0; i < N; i++ {
		fmt.Scan(&weights[i])
	}

	minWeight := weights[0]
	maxWeight := weights[0]

	for i := 1; i < N; i++ {
		if weights[i] < minWeight {
			minWeight = weights[i]
		}
		if weights[i] > maxWeight {
			maxWeight = weights[i]
		}
	}

	fmt.Printf("Berat terkecil: %.2f\n", minWeight)
	fmt.Printf("Berat terbesar: %.2f\n", maxWeight)
}
