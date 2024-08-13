package main

// 闭包
func Closure(name string) func() string {
	return func() string {
		return "hello" + name
	}
}

func Closure1(name string) func() int {
	var age = 0
	return func() int {
		age++
		return age
	}
}
