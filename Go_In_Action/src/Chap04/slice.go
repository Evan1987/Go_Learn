package main

import "fmt"

func test1() {
	// 赋值切片

	slice := make([] string, 5, 10)
	fmt.Println(cap(slice), len(slice))

	slice1 := []string{"red", "blue", "green", "yellow", "pink"}
	fmt.Println(cap(slice1), len(slice1))

	slice2 := []string{3: "a", 5: "b"}
	fmt.Println(slice2)
	fmt.Println(cap(slice2), len(slice2))
}

func test2() {
	// 切片增长

	slice := []int{10, 20, 30, 40, 50}  // cap = 5
	newSlice := slice[1:3]  // [i : j]
	fmt.Println(newSlice)  // [20, 30]
	fmt.Println(cap(newSlice), len(newSlice))  // 4, 2, cap = 5 - i

	slice[2] += 2
	fmt.Println(newSlice)  // [20, 32] 底层数据共享

	newSlice = append(newSlice, 60)  // [20, 32, 60]
	fmt.Println(slice)  // [10, 20, 32, 60, 50] 对应数据也被修改了
	fmt.Println(cap(newSlice), len(newSlice))  // 4, 3

	newSlice = append(newSlice, 70, 80)  // 一次增加两个元素， 会使长度超过原有容量，此时newSlice被复制到新的数组中
	fmt.Println(slice)  // [10, 20, 32, 60, 50] 并没有改变，此时已经与newSlice不共享了
	fmt.Println(cap(newSlice), len(newSlice))  // 8, 5 容量增加一倍，但不总是倍乘，容量超过1000时，0.25

}

func test3() {
	// 创建切片的三索引
	// 这样可以限制切片的容量，使其在append时可以创建新的底层数组，而不影响原始数据

	source := []string{"apple", "orange", "plum", "banana", "grape"}

	slice := source[2:3:4]  // [i:j:k]
	fmt.Println(slice)  // [plum]   source[i:j]
	fmt.Println(cap(slice), len(slice))  // 2, 1   k - i, j - i

	slice1 := source[2:4:4]  // cap 2, len 2
	slice1 = append(slice1, "pear")
	fmt.Println(source)  // no change
	fmt.Println(cap(slice1), len(slice1))  // 4, 3

	// 切片追加
	s1 := []int{1, 2}
	s2 := []int{3, 4}
	fmt.Println(append(s1, s2...))  // 注意"..."用法
}

func test4() {
	// 迭代切片

	slice := []int{10, 20, 30, 40}
	for index, value := range slice {  // range只提供每个元素的副本，并不是直接引用
		fmt.Printf("Index: %d\tValue: %d\n", index, value)
	}

	for index, value := range slice {
		fmt.Printf("Value: %d  Value-Addr: %X  Elem-Addr: %X\n", value, &value, &slice[index])
	}
	//Value: 10  Value-Addr: C00007A058  Elem-Addr: C000096000
	//Value: 20  Value-Addr: C00007A058  Elem-Addr: C000096008
	//Value: 30  Value-Addr: C00007A058  Elem-Addr: C000096010
	//Value: 40  Value-Addr: C00007A058  Elem-Addr: C000096018

	// 使用传统循环进行迭代控制
	for index := 1; index < len(slice); index += 2 {
		fmt.Printf("Index: %d  Value: %d\n", index, slice[index])
	}
}



func main() {
	test4()
}
