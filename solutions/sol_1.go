package main

import (
	"fmt"
)

func main() {
	var list = []int{10, 15, 3, 7}
	fmt.Println(bruteFind(list, 4, 17))

	fmt.Println(optimumFind(list, 4, 17))
}

//its pretty straight foreward!
func bruteFind(list []int, size int, k int) bool {
	for i := 0; i < size; i++ {
		for j := i; j < size; j++ {
			if list[i]+list[j] == k {
				return true
			}
		}
	}
	return false
}

/*
if given array is [10,15,3,7]
differenceMap will contain the difference of each number in array
eg. if k is 17
=> diff = 17 - array[0]
=> diff = 17 - 10
=> diff = 7

=> differenceMap[7] = true

next when we find 7 in array, we check differenceMap
to find if the number is present.

if yes the sum is present
*/
func optimumFind(list []int, size int, k int) bool {
	var differenceMap = make(map[int]bool)

	for i := 0; i < size; i++ {
		currentNum := list[i]

		if differenceMap[currentNum] {
			return true
		}
		diff := k - currentNum
		differenceMap[diff] = true
	}
	return false
}
