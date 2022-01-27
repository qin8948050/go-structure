package main

import (
	"fmt"
)

func main() {
	l1 := []int{4, 1, 2, 3, 5}
	result := BubbleSort(l1)
	fmt.Println(result)
}

func BubbleSort(l []int) []int {
	for i := 0; i <= len(l)-1; i++ {
		for j := 0; j < len(l)-i-1; j++ {
			if l[j] < l[j+1] {
				tmp := l[j]
				l[j] = l[j+1]
				l[j+1] = tmp
			}
		}
	}
	return l
}

func BinarySearch(list []int, searchValue int) int {
	low := 0
	high := len(list) - 1
	var mid int
	for low <= high {
		mid := low + (high-low)/2
		if list[mid] == searchValue {
			return mid
		} else if list[mid] > searchValue {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return mid
}
