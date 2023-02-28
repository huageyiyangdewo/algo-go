package sort

func QuickSort(arr []int) []int {
	lo := 0
	hi := len(arr) - 1
	sort(arr, lo, hi)

	return arr
}

func sort(arr []int, lo, hi int) {
	// 子数组长度为 1 时终止递归
	if lo >= hi {
		return
	}
	//对arr数组中，从lo到hi的元素进行切分
	pivot := partition(arr, lo, hi)
	//对左边分组中的元素进行排序
	//对右边分组中的元素进行排序
	sort(arr, lo, pivot-1)
	sort(arr, pivot+1, hi)
}

func partition(arr []int, lo, hi int) int {
	i, j := lo, hi
	for i < j {
		for i < j && arr[j] >= arr[lo] {
			j-- //从右向左找到首个小于基准数的值 或者 i < j 不满足为止
		}
		for i < j && arr[i] <= arr[lo] {
			i++ //从左向右找到首个大于基准数的值 或者 i < j 不满足为止
		}

		arr[i], arr[j] = arr[j], arr[i]
	}
	// 将基准数交换至两子数组的分界线
	arr[i], arr[lo] = arr[lo], arr[i]
	return i
}
