package testfunction

import "testing"

func Test_function1(t *testing.T) {
	type args struct {
		x []float64
	}
	tests := []struct {
		name  string
		args  args
		wantS float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := function1(tt.args.x); gotS != tt.wantS {
				t.Errorf("function1() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}
