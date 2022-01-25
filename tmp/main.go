package main

import (
	"fmt"
)

func main()  {
	l1:=[]int{1,2,3,4,5}
	result:=BinarySearch(l1,5)
	fmt.Println(result)
}


func BubbleSort(list []int) []int{
	for i:=0;i<=len(list)-1;i++{
		for j:=i;j<=len(list)-1;j++{
			if list[i]>list[j]{
				tmp:=list[i]
				list[i]=list[j]
				list[j]=tmp
			}
		}
	}
	return list
}

func BinarySearch(list []int,searchValue int) int{
	low:=0
	high:=len(list)-1
	var mid int
	for low<=high{
		mid:=low+(high-low)/2
		if list[mid]==searchValue{
			return mid
		} else if list[mid]>searchValue{
			high=mid-1
		} else {
			low=mid+1
		} 
	}
	return mid
}