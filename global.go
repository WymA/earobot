package earobot

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
