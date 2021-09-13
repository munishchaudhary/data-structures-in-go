package main

import (
	"fmt"
)

func main() {
	arr := []int {-10,-5,-3,-1,2,5,6,7,8,9,10,14,15,19,21,26,44,45,60}
	keys := []int {-5, 45, 61,-11}
	idx := 0
	is_available := false
	for _, key :=range keys {
		fmt.Printf("searching key %d in array %v \n", key, arr)
		is_available, idx = binarySerach(arr, key)
		if is_available {
			fmt.Printf("key %d found in array %v at index %d \n\n", key, arr, idx)
		} else {
			fmt.Printf("key %d not found in array%v \n\n", key, arr)
		}
	}
}

func binarySerach(arr []int,key int) (bool, int) {
	idx := -1
	is_available := false
	i := 0
	j := len(arr)-1
	mid := i+j/2
	for i < j {
		mid = (i+j)/2
		if key == arr[mid] {
			return true, mid
		} else if key > arr[mid] {
			i = mid+1
		} else {
			j = mid-1
		}
	}
	return is_available, idx
}

