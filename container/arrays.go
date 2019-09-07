package container

import "fmt"

func main() {
	var arr [4]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int

	fmt.Println(arr, arr2, arr3)
	fmt.Println(grid)

	for _, v := range arr3 {
		fmt.Println(v)
	}
}

func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
