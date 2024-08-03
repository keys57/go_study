package main

func Bool() {
	var a bool = true
	var b bool = false
	var c bool = (a && b)
	println(a || b)
	println(c)
	var d bool = !(a && b)
	println(d)
}
