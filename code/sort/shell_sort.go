package sort

//ShellSort  排序原理：
// 1.选定一个增长量h，按照增长量h作为数据分组的依据，对数据进行分组；
// 2.对分好组的每一组数据完成插入排序；
// 3.减小增长量，最小减为1，重复第二步操作。
// 希尔排序是插入排序的一种，又称“缩小增量排序”，是插入排序算法的一种更高效的改进版本。
func ShellSort(arr []int) []int {
	N := len(arr)
	var h int
	// 确定增长量h的最大值
	for h = 1; h < (N/2); h++ {
		h = h * 2 + 1
	}

	//当增长量h小于1，排序结束
	for h >= 1 {
		//找到待插入的元素
		for i := h; i < N; i++ {
			//arr[i]就是待插入的元素
			//把arr[i]插入到arr[i-h],arr[i-2h],arr[i-3h]...序列中
			for j := i; j >=h ; j-=h {
				//arr[j]就是待插入元素，依次和arr[j-h],arr[j-2h],arr[j-3h]进行比较，
				//如果arr[j]小，那么交换位置，如果不小于，arr[j]大，则插入完成。
				if arr[j-h] > arr[j] {
					arr[j-h], arr[j] = arr[j], arr[j-h]
				} else {
					break
				}
			}
		}

		h /= 2
	}

	return arr
}
