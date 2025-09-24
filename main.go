package main

import (
	"fmt"
	"github.com/Frame87/module03/utils"
)

func main() {
	stringSlice := []string{"Hello", "World", "Go"}
	fmt.Println("Contains 'World':", utils.Contains(stringSlice, "World"))

	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("ContainsInt 3:", utils.ContainsInt(intSlice, 3))
	fmt.Println("InSliceInt 3:", utils.InSliceInt(intSlice, 3))
}
