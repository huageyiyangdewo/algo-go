package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type test struct {
		input []int
		want []int
	}

	tests := map[string]test{
		"the first test case": {input: []int{3,4,5,7,8,1,7}, want: []int{1,3,4,5,7,7,8}},
		"the second test case": {input: []int{9,6,4,2,1}, want: []int{1,2,4,6,9}},
		"the third test case": {input: []int{-1,1,10,99,43,992,-2}, want: []int{-2,-1,1,10,43,99,992}},
		"the fourth test case": {input: []int{6,0,3,-99,-77,88}, want: []int{-99,-77,0,3,6,88}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := QuickSort(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected:%#v, got:%#v", tc.want, got)
			}
		})
	}
}
