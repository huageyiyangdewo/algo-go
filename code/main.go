package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	generateShellSortData()

}

func reverse(arr []int) []int {
	n := len(arr)
	var tmp int

	end := n - 1
	for start := 0; start <= end; start++ {
		tmp = arr[start]
		arr[start] = arr[end]
		arr[end] = tmp
		end--
	}

	return arr
}

func reverse2(arr []int) []int {
	n := len(arr)
	tmp := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		tmp[n-1-i] = arr[i]
	}

	return tmp
}

//生成1000000-0之间的数据存放在sort文件夹下面的 reverse_shell_sort.txt 里面
func generateShellSortData()  {
	err := createDir("./sort")
	if err != nil {
		return
	}

	err = createFile("./sort/reverse_shell_sort.txt")
	if err != nil {
		fmt.Println("文件创建失败", err)
		return
	}
}

//hasDir 判断文件或者文件夹是否存在
func hasDir(path string) (bool, error) {
	_, err := os.Stat(path) // 获取 文件或文件夹信息
	if err == nil { // 文件或者文件夹存在
		return true, nil
	}

	//判断错误类型
	if os.IsNotExist(err) {
		//不存在
		return false, nil
	}

	return false, nil
}

func createDir(path string) error {
	exist, err := hasDir(path)
	if err != nil {
		fmt.Printf("获取文件夹异常：%v\n", err)
		return err
	}

	if exist {
		fmt.Println("文件夹已经存在")
		return nil
	}

	err = os.Mkdir(path, 0666)
	if err != nil {
		fmt.Printf("创建文件夹异常：%v\n", err)
		return err
	} else {
		fmt.Println("文件夹创建成功")
	}

	return nil
}

func createFile(path string) error {
	exist, err := hasDir(path)
	if err != nil {
		fmt.Printf("获取文件异常：%v\n", err)
		return err
	}

	if exist {
		fmt.Println("文件已经存在")
		return nil
	}

	_, err = os.Create(path)
	if err != nil {
		fmt.Printf("创建文件异常：%v\n", err)
		return err
	} else {
		fmt.Println("文件创建成功!!")
		f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0666)
		if err != nil {
			fmt.Printf("打开文件：%v 失败，原因：%v\n", path, err)
			return err
		}

		var tmp strings.Builder
		count := 0
		for i := 6000; i > 0; i-- {
			tmp.WriteString(strconv.Itoa(i))
			if i == 1 {

			 } else {
				tmp.WriteString(",")
			}

			if count == 1000 {
				_, _ = f.WriteString(tmp.String())
				count = 0
				tmp.Reset()
			}
			count++

			if i == 1 {
				_, _ = f.WriteString(tmp.String())
			}
		}
	}

	return nil
}
