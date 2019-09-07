package container

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := arr[2:6]

	fmt.Println(s)
	s = s[3:5]
	fmt.Println(s)
	fmt.Println(s, len(s), cap(s))
	//fmt.Println(s[3]) // index out of range
	s = s[:4] // Extending slice
	fmt.Println(s)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 10)
	fmt.Println(s, len(s), cap(s))
	s = append(s, 11)
	fmt.Println(s, len(s), cap(s))
	s = append(s, 12)
	fmt.Println(s, len(s), cap(s))
	s = append(s, 13)
	s = append(s, 13)
	s = append(s, 13)
	s = append(s, 13)
	s = append(s, 13)
	fmt.Println(s, len(s), cap(s))

	newSlice()
}

func updateSlice(s []int) {

}

func newSlice() {
	var s []int
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s2 := []int{2, 4, 6, 8}
	printSlice(s2)

	s3 := make([]int, 16)
	s4 := make([]int, 10, 32) // len: 10, cap: 32
	printSlice(s3)
	printSlice(s4)

	copy(s3, s2) // copy: dst, src
	printSlice(s3)

	s3 = append(s3[:3], s3[4:]...) // delete
	printSlice(s3)

	// pop from front
	front := s3[0]
	s3 = s3[1:]
	printSlice(s3)
	// pop from back
	back := s3[len(s3)-1]
	s3 = s3[:len(s3)-1]
	printSlice(s3)

	fmt.Println("front back", front, back)

}

func printSlice(s []int) {
	fmt.Printf("s=%v, len(s)=%d, cap(s)=%d\n", s, len(s), cap(s))
}
