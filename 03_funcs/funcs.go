package main

// Func1没有任何参数
func Func1() {

}

// Func2有一个参数
func Func2(a int) {

}

// Func3有多个参数,一个类型
func Func3(a, b int) {

}

func Func4(a, b int, c string) {

}

func Func5(a, b int, c string) string {
	//有返回值要保证一定返回
	return "hello world"
}

// 多个返回值
func Func6(a, b int, c string) (string, string) {
	//有返回值要保证一定返回
	return "hello world", "test2"
}

// 返回值带名称
func Func7() (name string, age int) {
	return "test", 18
}

func Func8() (name string, age int) {
	//等价于 "",0
	//对应类型的默认值
	return
}
