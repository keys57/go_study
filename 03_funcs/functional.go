package main

func Functional4() {
	println("Hello World")
}

func Functional5(age int) {

}

func UseFunctional4() {
	myFunc := Functional4
	myFunc()

	myFunc5 := Functional5
	myFunc5(18)

}

func Functional6() {
	//新定义了一个方法，赋值给了 fn
	fn := func() string {
		return "Hello"
	}
	fn()
}

// 意思是返回一个，返回一个string 的无参数方法
func Functional7() func() string {
	return func() string {
		return "Hello"
	}
}

// 匿名方法立刻发起调用
func Functional8() {
	//新定义了一个方法，赋值给了 fn
	fn := func() string {
		return "Hello"
	}()
	println(fn)

}
