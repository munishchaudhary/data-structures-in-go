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
	mergeSort(arr, left, right)
	fmt.Println("sorted array values are", arr)
}

func mergeSort(arr []int,left int, right int) {
	mid := (left+right)/2
	if left < right {
		mergeSort(arr, left, mid)
		mergeSort(arr, mid+1, right)
		mergeArray(arr, left, mid, right)
	}
}

func mergeArray(arr []int, left int, mid int, right int) {
	i := left
	j := mid+1
	k := left
	temp := make([]int, right+1)

	for i <= mid && j <=right {
		if arr[i] < arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
	}
	if i > mid {
		for j<=right {
			temp[k] = arr[j]
			j++
			k++
		}
	} else if j > right {
		for i<=mid {
			temp[k] = arr[i]
			i++
			k++
		}
	}
	for i:=left; i<=right; i++{
		arr[i] = temp[i]
	}
}



