package main

import (
	"fmt"
)

func main() {
	arr := []int {10, 5, 10, 5, 55, 22, 66, 77, 0, 4, 7, 8, 1, 5, -1, -10}
	fmt.Println("unsorted array values are", arr)
	size := len(arr)
	left := 0
	right := size-1
	quickSort(arr, left, right)
	fmt.Println("sorted array values are", arr)
}

func quickSort(arr []int, left int, right int) {
	if left < right {
		pivot := partition(arr, left, right)
		quickSort(arr, left, pivot-1)
		quickSort(arr, pivot+1, right)
	}
}

func partition(arr []int,left int, right int) int {
	pivot := arr[left]
	i := left
	j := right
	temp := 0
	for i < j {
		for  i<=right && arr[i] <= pivot {
			i++
		}
		for j>=left && arr[j] > pivot{
			j--
		}
		if i < j {
			temp = arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}
	temp = arr[left]
	arr[left] = arr[j]
	arr[j] = temp
	return j
}

