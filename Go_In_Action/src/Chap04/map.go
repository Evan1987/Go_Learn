package main

import "fmt"


func map_test() {
	// 映射创建与使用

	dict := make(map[string]string)

	dict = map[string]string {"Red": "#da1337", "Orange": "#e95a32"}
	fmt.Println(dict)

	value, exists := dict["Red"]
	fmt.Println(value, exists)  // "#da1337" true

	value, exists = dict["Green"]
	fmt.Println(value, exists)  // false
	fmt.Println(len(value))  // 0

	delete(dict, "Red")  // inplaced, not return

	for key, value := range dict {
		fmt.Printf("Key: %s  Value: %s\n", key, value)
	}
}




func main() {
	map_test()
}

