package main

func Defer() {
	defer func() {
		println("第一个defer")
	}()

	defer func() {
		println("第二个defer")
	}()
}

func DeferClosure() {
	i := 0
	defer func() {
		println(i)
	}()

	i = 1
}
