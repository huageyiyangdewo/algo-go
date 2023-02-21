package sort


// BubbleSort 冒泡排序
// 排序原理：
// 1. 比较相邻的元素。如果前一个元素比后一个元素大，就交换这两个元素的位置。
// 2. 对每一对相邻元素做同样的工作，从开始第一对元素到结尾的最后一对元素。最终最后位置的元素
// 就是最大值。
func BubbleSort(arr []int) []int {
	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}
