package main

import "fmt"

func Byte() {
	var a byte = 'a'
	println(a)
	println(fmt.Sprintf("%c", a))

	var str string = "this is a string"
	var bs []byte = []byte(str)
	println(str, bs)
}
