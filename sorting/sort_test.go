package sort

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	arr []int
}

type test struct {
	name string
	args args
}

var tests = []test{
	{
		args: args{
			arr: []int{5, 3, 6, 2, 4, 1},
		},
	},
	{
		args: args{
			arr: []int{100, 2000, -1, -59, 67, 342, 897, 374, 85536},
		},
	},
	{
		args: args{
			arr: []int{},
		},
	},
	{
		args: args{
			arr: []int{-1, -1, -2, -100, -2334, -43434},
		},
	},
	{
		args: args{
			arr: []int{0},
		},
	},
	{
		args: args{
			arr: []int{1, 2, 3, 4, 5},
		},
	},
}

func Test_bubbleSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := make([]int, len(tt.args.arr))
			copy(want, tt.args.arr)
			sort.Ints(want)
			bubbleSort(tt.args.arr)
			assert.Equal(t, want, tt.args.arr)
			t.Logf("Sorted Arr :: %+v", tt.args.arr)
		})
	}
}

func Test_insertionSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := make([]int, len(tt.args.arr))
			copy(want, tt.args.arr)
			sort.Ints(want)
			insertionSort(tt.args.arr)
			assert.Equal(t, want, tt.args.arr)
			t.Logf("Sorted Arr :: %+v", tt.args.arr)
		})
	}
}

func Test_selectionSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := make([]int, len(tt.args.arr))
			copy(want, tt.args.arr)
			sort.Ints(want)
			selectionSort(tt.args.arr)
			assert.Equal(t, want, tt.args.arr)
			t.Logf("Sorted Arr :: %+v", tt.args.arr)
		})
	}
}

func Test_mergeSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := make([]int, len(tt.args.arr))
			copy(want, tt.args.arr)
			sort.Ints(want)
			mergeSort(tt.args.arr)
			assert.Equal(t, want, tt.args.arr)
			t.Logf("Sorted Arr :: %+v", tt.args.arr)
		})
	}
}

func Test_quickSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := make([]int, len(tt.args.arr))
			copy(want, tt.args.arr)
			sort.Ints(want)
			quickSort(tt.args.arr)
			assert.Equal(t, want, tt.args.arr)
			t.Logf("Sorted Arr :: %+v", tt.args.arr)
		})
	}
}
