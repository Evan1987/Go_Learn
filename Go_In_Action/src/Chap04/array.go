package main

import "fmt"


func main() {
	var arr1 [5]int

	arr2 := [5]int{10, 20, 30, 40, 50}

	arr3 := []int{10, 20, 30, 40, 50}

	arr4 := [5]int{1: 10, 2: 20}  // [0 10 20 0 0]

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)

	arr5 := [5] *int {0: new(int)}
	fmt.Println(arr5)

	fmt.Println(*arr5[0])  // 0
	*arr5[0] = 10

}
