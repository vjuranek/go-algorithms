package main

import (
	"fmt"
)

func restoreHeap(arr []int, begin int, end int) []int {
	i := arr[begin-1]
	var j int
	for k := begin; k <= end/2; {
		j = 2 * k // right child (2(k-1)+2 == 2k)
		if (j < end) && (arr[j-1] < arr[j]) { // compare left (2(k-1)+1 == 2k-1) and right child
			j++
		}
		if i >= arr[j-1] { //compare parent and bigger one of its children
			break
		} else { //switch (original) parent and bigger one of its children
			arr[k-1] = arr[j-1]
			k = j
			arr[k-1] = i
		}
	}
	return arr
}

func heapsort(arr []int) []int {
	length := len(arr)
	for i := length / 2; i > 0; i-- {
		arr = restoreHeap(arr, i, length)
	}

	for i := length; i > 1; i-- {
		tmp := arr[0]
		arr[0] = arr[i-1]
		arr[i-1] = tmp
		arr = restoreHeap(arr, 1, i-1)
	}

	return arr
}

func main() {
	arr := []int{2, 5, 1, 3, 4}
	fmt.Println(arr)
	arr = heapsort(arr)
	fmt.Println(arr)
}
