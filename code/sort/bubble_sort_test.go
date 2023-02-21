package sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T)  {

	// 定义一个测试用例类型
	type test struct {
		input []int
		want []int
	}

	tests := []test{
		{input: []int{3,4,5,7,8,1,7}, want: []int{1,3,4,5,7,7,8}},
		{input: []int{9,6,4,2,1}, want: []int{1,2,4,6,9}},
		{input: []int{-1,1,10,99,43,992,-2}, want: []int{-2,-1,1,10,43,99,992}},
		{input: []int{6,0,3,-99,-77,88}, want: []int{-99,-77,0,3,6,88}},
	}

	for _, tc := range tests {
		got := BubbleSort(tc.input)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("expected: %#v, got: %#v", tc.want, got)
		}
	}
}
