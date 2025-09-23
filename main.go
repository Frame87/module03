package main

import (
	"fmt"
	"github.com/Frame87/module03/utils"
)

func main() {
	// Использование оригинальной версии
	fmt.Println("Original utils version:")
	stringSlice := []string{"Hello", "World", "Go"}
	fmt.Println("Contains 'World':", utils.Contains(stringSlice, "World"))

	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("ContainsInt 3:", utils.ContainsInt(intSlice, 3))

	// Использование второй функции в том же пакете
	fmt.Println("\nUtils v2 version:")
	fmt.Println("InSliceInt 3:", utils.InSliceInt(intSlice, 3))
	fmt.Println("InSliceInt 6:", utils.InSliceInt(intSlice, 6))
}
