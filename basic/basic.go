package basic

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
	"strconv"
)

func main() {
	fmt.Println("Hello world")

	euler()
	consts()

	fmt.Println(
		grade(59),
		grade(60),
		grade(82),
		grade(99),
		grade(100),
	)

	fmt.Println(convertToBin(123456789))
	fmt.Println(apply(3, 4, pow))
	fmt.Println(apply(3, 4, func(a int, b int) int {
		return a + b
	}))
	fmt.Println(sumArgs(1, 2, 3, 4, 5, 6, 7, 8, 9))

	pointer()
}

func euler() {
	//fmt.Printf("%.3f", math.Pow(math.E, 1i*math.Pi) + 1)
}

func consts() {
	const (
		c = iota
		java
		python
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(c, java, python)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func grade(score int) string {
	g := ""
	switch {
	case score < 60 && score >= 0:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		panic(fmt.Sprintf("Wrong score: %d", score))
	}
	return g
}

func convertToBin(i int) string {
	result := ""
	for ; i > 0; i /= 2 {
		lsb := i % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported op: %s", op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func apply(a, b int, op func(int, int) int) int {
	p := reflect.ValueOf(op).Pointer()
	funcName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (a=%d, b=%d)", funcName, a, b)
	return op(a, b)
}

func sumArgs(args ...int) int {
	sum := 0
	for _, i := range args {
		sum += i
	}
	return sum
}

func pointer() {
	var a int = 1
	var pa *int = &a
	*pa = 99
	fmt.Printf("Value of a: %d", a)
}

func passByVal(a int) {

}

func passByRef(a *int) {

}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func swapAB(a, b int) (int, int) {
	return b, a
}
