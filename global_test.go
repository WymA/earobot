package earobot

import (
	"reflect"
	"testing"
)

func TestDistVector(t *testing.T) {
	type args struct {
		vectors1 []float64
		vectors2 []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistVector(tt.args.vectors1, tt.args.vectors2); got != tt.want {
				t.Errorf("DistVector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinFastSort(t *testing.T) {
	type args struct {
		x   []float64
		idx []int
		n   int
		m   int
	}
	tests := []struct {
		name  string
		args  args
		want  []float64
		want1 []int
	}{
		{
			"1",
			args{
				x:   []float64{2, 1, 5, 3, 4},
				idx: []int{0, 1, 2, 3, 4},
				n:   0,
				m:   5,
			},
			[]float64{1, 2, 3, 4, 5},
			[]int{0, 1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MinFastSort(tt.args.x, tt.args.idx, tt.args.n, tt.args.m)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinFastSort() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("MinFastSort() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
