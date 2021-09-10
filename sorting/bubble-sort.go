package main

import (
	"fmt"
)

func main() {
	arr := []int {10, 5, 10, 5, 55, 22, 66, 77, 0, 4, 7, 8, 1, 5, -1, -10} 
	size := len(arr)
	temp := 0
	swapped := false
	fmt.Println("unsorted array values are", arr)
	for i:=0; i<size-1; i++ {
		swapped = false
		for j:=0; j<(size-i)-1; j++ {
			if arr[j] > arr[j+1] {
			        swapped = true
				temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
		// if swapped is false means array is now sorted, no need to process it further 
		if !swapped {
			break
		}
	}
	fmt.Println("sorted array values are", arr)
}

