package main

import (
	"fmt"
)

func main() {
	arr := []int {10, 5, 10, 5, 55, 22, 66, 77, 0, 4, 7, 8, 1, 5, -1, -10 ,-100}
	fmt.Println("unsorted array values are", arr)
	selectionSort(arr)
	fmt.Println("sorted array values are", arr)
}

func selectionSort(arr []int) {
        temp := 0
        min_idx := 0
        size := len(arr)
        for i:=0; i<size-1; i++ {
                min_idx = i
                for j := i+1; j<size; j++ {
                        if arr[j] < arr[min_idx] {
                                min_idx = j
                        }
                }
                if min_idx != i {
                        temp = arr[i]
                        arr[i] = arr[min_idx]
                        arr[min_idx] = temp
                }
        }
}
