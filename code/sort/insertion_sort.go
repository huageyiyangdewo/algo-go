package sort

//InsertionSort 插入排序
//排序原理：
//1.把所有的元素分为两组，已经排序的和未排序的；
//2.找到未排序的组中的第一个元素，向已经排序的组中进行插入；
//3.倒叙遍历已经排序的元素，依次和待插入的元素进行比较，直到找到一个元素小于等于待插入元素，
//那么就把待插入元素放到这个位置，其他的元素向后移动一位；
func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {  //相当于未排序的数组
		for j := i; j > 0; j-- {  //相当于已排序的数组，倒序遍历，更加简单
			if arr[j] <= arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}

	return arr
}
