package reserverlinkedlist206

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}},
			want: &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}}}}},
		},
		{
			head: &ListNode{Val: 1, Next: nil},
			want: &ListNode{Val: 1, Next: nil},
		},
		{
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}},
			want: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
