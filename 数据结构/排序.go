package main

import (
	"fmt"
	"math/rand"
	"time"
)

func BubbleSort(arr []int) { //冒泡排序：最前两个数开始比较，谁大谁往右，再依次往后推。每循环一次n-1
	n := len(arr)
	for i := 0; i < n-1; i++ { // 外层循环控制排序轮数
		for j := 0; j < n-i-1; j++ { // 内层循环控制每一轮比较次数
			if arr[j] > arr[j+1] { // 比较相邻两个元素的大小，如果前一个比后一个大则交换位置
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func SelectionSort(arr []int) { //选择排序：先遍历整个数组选出最小数，排在最前，然后从未排序区间取最小值
	n := len(arr)
	for i := 0; i < n-1; i++ { // 外层循环控制排序轮数
		minIndex := i                // 记录最小值下标
		for j := i + 1; j < n; j++ { // 查找未排序区间内的最小值
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i] // 将最小值与未排序区间的第一个元素交换位置
	}
}

func InsertionSort(arr []int) { //插入排序:从左选两个数，依次向右遍历，将遍历好的数往左放
	n := len(arr)
	for i := 1; i < n; i++ { // 外层循环控制插入次数
		key := arr[i] // 缓存当前需要插入的元素
		j := i - 1
		for ; j >= 0 && arr[j] > key; j-- { // 内层循环控制比较次数
			arr[j+1] = arr[j] // 比较大小并将元素向右移动
		}
		arr[j+1] = key // 插入元素
	}
}

func QuickSort(arr []int, left, right int) { // QuickSort 快速排序算法实现
	if left < right {
		pivot := partition(arr, left, right) // 分区
		QuickSort(arr, left, pivot-1)        // 递归排序左右子数组
		QuickSort(arr, pivot+1, right)
	}
}

func partition(arr []int, left, right int) int { // partition 分区函数，用于快速排序
	pivot := arr[right]             // 取最后一个元素作为基准值
	i := left - 1                   // 定义初始分割点
	for j := left; j < right; j++ { // 循环遍历区间内的所有元素
		if arr[j] <= pivot { // 如果当前元素小于等于基准值，则将i右移并交换i和j的位置
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[right] = arr[right], arr[i+1] // 将i后面的元素与基准值交换位置
	return i + 1                                // 返回分割点
}

func MergeSort(arr []int) []int { // MergeSort 归并排序算法实现
	if len(arr) == 1 { // 判断数组长度是否为1
		return arr
	}
	mid := len(arr) / 2          // 求中间值
	left := MergeSort(arr[:mid]) // 递归分别对左右子数组进行排序
	right := MergeSort(arr[mid:])

	return merge(left, right) // 合并左右子数组
}

func merge(left, right []int) []int { // merge 辅助函数，用于合并两个有序数组
	result := make([]int, 0, len(left)+len(right)) // 初始化结果数组
	i, j := 0, 0                                   // 定义左右指针
	for i < len(left) && j < len(right) {          // 循环比较左右数组中的元素，并将较小的元素加入结果数组中
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...) // 将剩余的元素加入结果数组中
	result = append(result, right[j:]...)
	return result
}

func main() {
	rand.NewSource(time.Now().UnixNano()) // 生成随机数切片
	arr := make([]int, 10)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	fmt.Println("原始数据：", arr)

	BubbleSort(arr) // 冒泡排序
	fmt.Println("冒泡排序：", arr)

	SelectionSort(arr) // 选择排序
	fmt.Println("选择排序：", arr)

	InsertionSort(arr) // 插入排序
	fmt.Println("插入排序：", arr)

	QuickSort(arr, 0, len(arr)-1) // 快速排序
	fmt.Println("快速排序：", arr)

	arr = MergeSort(arr) // 归并排序
	fmt.Println("归并排序：", arr)
}
