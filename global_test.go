package earobot

import "testing"

func TestMinFastSort(t *testing.T) {
	type args struct {
		x   []float64
		idx []int
		n   int
		m   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MinFastSort(tt.args.x, tt.args.idx, tt.args.n, tt.args.m)
		})
	}
}
