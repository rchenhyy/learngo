package container

import "fmt"

func main() {
	m := map[string]string{
		"a": "b",
		"c": "d",
		"e": "f",
	}
	m2 := make(map[string]string) // empty map
	var m3 map[string]string      // nil

	fmt.Println(m, m2, m3)
	v := m["x"] // zero value
	fmt.Printf("m[\"x\"]=%q\n", v)

	v, ok := m["x"]
	if ok {
		fmt.Printf("m[\"x\"]=%q\n", v)
	} else {
		fmt.Printf("m[\"x\"] not exists\n")
	}

	delete(m, "x")
	delete(m, "c")
	fmt.Printf("m[\"c\"]=%q\n", m["c"])
}


func lenOfNoneRepeatingSubStr(s string) int {
	return 0
}