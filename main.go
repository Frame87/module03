package main

import (
	"fmt"
	"github.com/Frame87/module03/utils"
	utilsV2 "github.com/Frame87/module03/v2"
)

func main() {
	// Использование оригинальной версии
	fmt.Println("Original utils version:")
	stringSlice := []string{"Hello", "World", "Go"}
	fmt.Println("Contains 'World':", utils.Contains(stringSlice, "World"))

	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("ContainsInt 3:", utils.ContainsInt(intSlice, 3))

	// Использование второй версии
	fmt.Println("\nUtils v2 version:")
	fmt.Println("InSliceInt 3:", utilsV2.InSliceInt(intSlice, 3))
	fmt.Println("InSliceInt 6:", utilsV2.InSliceInt(intSlice, 6))
}
