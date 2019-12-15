package plusone66

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	tests := []struct {
		digits []int
		want   []int
	}{
		{
			digits: []int{, 9},
			want:   []int{1, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := plusOne(tt.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
