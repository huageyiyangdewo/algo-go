package sort

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestMergeSort(t *testing.T) {
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
			got := MergeSort(tc.input)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected:%#v, got:%#v", tc.want, got)
			}
		})
	}

	content, err := ioutil.ReadFile("./reverse_shell_sort.txt")
	if err != nil {
		fmt.Printf("read file error, %v\n", err)
		return
	}

	tmp := strings.Split(string(content), ",")
	var sortArr []int
	for _, value := range tmp {
		i, _ := strconv.Atoi(value)
		sortArr = append(sortArr, i)
	}

	sortArr2 := make([]int, len(sortArr), len(sortArr))
	copy(sortArr2, sortArr)
	t.Log(len(sortArr))
	t.Log(len(sortArr2))

	now := time.Now()
	ShellSort(sortArr)
	t.Log("希尔排序花费时间：", time.Now().Sub(now))

	now = time.Now()
	MergeSort(sortArr2)
	t.Log("归并排序花费时间：", time.Now().Sub(now))
}
