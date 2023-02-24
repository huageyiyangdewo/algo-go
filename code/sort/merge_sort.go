package sort

//MergeSort 排序原理：
// 1.尽可能的一组数据拆分成两个元素相等的子组，并对每一个子组继续拆分，直到拆分后的每个子组的元素个数是 1为止。
// 2.将相邻的两个子组进行合并成一个有序的大组；
// 3.不断的重复步骤2，直到最终只有一个组为止。
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	//获取分区位置 p
	p := len(arr) / 2

	left := MergeSort(arr[:p])
	right := MergeSort(arr[p:])

	return merge(left, right)
}

//merge 归并排序
func merge(left []int, right []int) []int {
	i, j := 0, 0
	m, n := len(left), len(right)

	//用于存放结果集
	var result []int
	for  {
		//任何一个区间先遍历完则退出
		if i >= m || j >= n {
			break
		}

		// 对所有区间数据进行排序
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	//下面两个循环，只会执行其中的一个
	// 如果左侧区间还没有遍历完，将剩余结果放到结果集
	if i != m {
		for ; i < m; i++ {
			result = append(result, left[i])
		}
	}
	// 如果右侧区间还没有遍历完，将剩余结果放到结果集
	if j != n {
		for ; j < n; j++ {
			result = append(result, right[j])
		}
	}

	//返回排序后的结果集
	return result
}
