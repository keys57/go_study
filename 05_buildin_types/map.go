package main

import "fmt"

func Map() {
	m1 := map[string]int{
		"a": 1,
	}
	m1["b"] = 2
	fmt.Println(m1)

	m2 := make(map[string]int, 12)
	m2["a"] = 1

	val, ok := m2["b"]
	if ok {
		println(val)
	}

	delete(m2, "a")
}
