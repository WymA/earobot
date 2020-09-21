package earobot

import "math"

// Minimal Fast Sort
func MinFastSort(x []float64, idx []int, n int, m int) {

	for i := 0; i < m; i++ {
		for j := i + 1; j < n; j++ {
			if x[i] > x[j] {
				x[i], x[j] = x[j], x[i]
				idx[i], idx[j] = idx[j], idx[i]
			}
		}
	}
}

// DistVector calculates the distance of two vector
func DistVector(vectors1 []float64, vectors2 []float64) float64 {

	sum := 0.0
	for index, vec1 := range vectors1 {
		sum += (vec1 - vectors2[index]) * (vec1 - vectors2[index])
	}

	return math.Sqrt(sum)
}

// ObservePointDistVector calculates the distance of two vector
func ObservePointDistVector(vectors1 []float64, vectors2 []float64) float64 {
	sum := 0.0
	for index, vec1 := range vectors1 {
		sum += (vec1 - vectors2[index]) * (vec1 - vectors2[index])
	}

	return math.Sqrt(sum)
}
