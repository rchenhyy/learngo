package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(incr int) int {
		sum += incr
		return sum
	}
}

type iAdder func(int) (int, iAdder) // 递归定义的无状态函数

func adder2(base int) iAdder {
	return func(i int) (int, iAdder) {
		rs := base + i
		return rs, adder2(rs) // rs：既是add的结果，又是下一次add的base
	};
}

func main2() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(a(i))
	}

	a2 := adder2(0)
	for i := 0; i < 10; i++ {
		var sum int
		sum, a2 = a2(i)
		fmt.Println(i, sum)
	}
}
