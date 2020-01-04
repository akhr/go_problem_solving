// To execute Go code, please declare a func main() in a package "main"

// Given Problem statement:
// Given a Natural number N > 0 find all its factors that are prime numbers

// Definition
// Divisors:
// Any number that divides another number. eg: Divisors of 12 are [1,2,3,4,6,12]
// Prime numbers:
// A prime number is a whole number greater than 1 whose only factors are 1 and itself. Numbers with exact two factors are prime numbers eg: Divisors of 11 are [1,11]
// Whole numbers are all natural numbers including 0 e.g. 0, 1, 2, 3, 4

// Input:
//         70
// Divisors
// [2,5,7,10,14,....35,70]
// Prime Divisors Output:
//         prime divisors of 70 : [2,5,7]

// Input:
//         13
// Output:
//         prime divisors of 13 : [13]

// Input:
//         1
// Output:
//         prime divisors of 1 : []

//21
// [3,7]
//14
//[2,7]

//78
//[2, 3, 13]

package prime

import (
	"reflect"
	"testing"
)

func Test_getPrimeFactors(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				n: 70,
			},
			want: []int{2, 5, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPrimeFactors(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPrimeFactors() = %v, want %v", got, tt.want)
			}
		})
	}
}
