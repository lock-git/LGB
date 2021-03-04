package main

import "fmt"

func main() {
	fmt.Println("Hello Word!")
	arr()
}

func arr() {
	arr1 := [5]string{"a", "b", "c", "d", "e"}
	for k, v := range arr1 {
		fmt.Printf("key:%d      value:%s\n", k, v)
	}
}
