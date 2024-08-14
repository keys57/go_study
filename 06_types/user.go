package main

import "fmt"

func NewUser() {
	//初始化结构体
	u := User{}
	fmt.Printf("%v \n", u)
	fmt.Printf("%+v \n", u)

	//up是一个指针
	up := &User{}
	fmt.Printf("%+v \n", up)
	up2 := new(User)
	fmt.Printf("%+v \n", up2)
}

type User struct {
	Name string
	Age  int
}
