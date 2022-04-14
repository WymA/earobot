package pathplanning

import (
	"testing"
)

func Test_allocate_pop(t *testing.T) {
	type args struct {
		pop    *Population
		size   int
		objNum int
	}
	tests := []struct {
		name        string
		args        args
		wantIndSize int
		wantObjSize int
	}{
		{
			"seze-100-obj-10",
			args{
				pop:    &Population{Individuals: []Individual{}},
				size:   100,
				objNum: 10,
			},
			100,
			10,
		},
		{
			"seze-1-obj-1",
			args{
				pop:    &Population{Individuals: []Individual{}},
				size:   1,
				objNum: 1,
			},
			1,
			1,
		},
		{
			"seze-0-obj-1",
			args{
				pop:    &Population{Individuals: []Individual{}},
				size:   0,
				objNum: 1,
			},
			0,
			0,
		},
		{
			"seze-1-obj-0",
			args{
				pop:    &Population{Individuals: []Individual{}},
				size:   1,
				objNum: 0,
			},
			1,
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIndSize, gotObjSize := allocate_pop(tt.args.pop, tt.args.size, tt.args.objNum)
			if gotIndSize != tt.wantIndSize {
				t.Errorf("allocate_pop() gotIndSize = %v, want %v", gotIndSize, tt.wantIndSize)
			}
			if gotObjSize != tt.wantObjSize {
				t.Errorf("allocate_pop() gotObjSize = %v, want %v", gotObjSize, tt.wantObjSize)
			}
		})
	}
}
