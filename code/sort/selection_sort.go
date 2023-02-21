package sort


// SelectionSort 选择排序
//排序原理：
//1.每一次遍历的过程中，都假定第一个索引处的元素是最小值，和其他索引处的值依次进行比较，
//如果当前索引处的值大于其他某个索引处的值，则假定其他某个索引出的值为最小值，
//最后可以找到最小值所在的索引
//2.交换第一个索引处和最小值所在的索引处的值
func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr) - 1; i++ {
		minIndex := i // 假定第一个索引处的元素是最小值
		for j := i + 1; j < len(arr); j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}

		// 交换第一个索引处和最小值所在的索引处的值
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

	return arr
}
