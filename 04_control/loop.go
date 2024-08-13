package main

func ForLoop() {
	for i := 0; i < 10; i++ {
		println(i)
	}
}

func Loop2() {
	i := 0
	for i < 10 {
		i++
		println(i)
	}

	for true {
		i++
		println(i)
	}
}

func ForArr() {
	arr := [3]int{1, 2, 3}
	for i, v := range arr {
		println(i, v)
	}
}
