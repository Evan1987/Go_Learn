package main

import (
	"fmt"
	"./words"
	"io/ioutil"
	"os"
)


func main() {
	// fileName 如果路径有空格，则空格需要转义
	fileName := os.Args[1]


	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	text := string(contents)

	count := words.CountWords(text)
	fmt.Printf("There are %d words in your text. \n", count)
}