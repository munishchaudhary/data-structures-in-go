package main

import (
	"fmt"
)

func main() {
	arr := []int {10, 5, 10, 5, 55, 22, 66, 77, 0, 4, 7, 8, 1, 5, -1, -10}
	fmt.Println("unsorted array values are", arr)
	insertionSort(arr)
	fmt.Println("sorted array values are", arr)
}

func insertionSort(arr []int) {
        size := len(arr)
        temp := 0
        j := 0
        for i:=1; i<size; i++ {
                temp = arr[i]
                for j = i-1; j>=0 && arr[j]>temp; j-- {
                        arr[j+1] = arr[j]
                }
                arr[j+1] = temp
        }
}
