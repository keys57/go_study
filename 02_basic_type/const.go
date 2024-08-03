package main

const External = "包外"
const internal = "包内"

func main() {
	const test1 = "测试1"
	const test2 = 1
}

const (
	statusA = iota
	statusB
	statusC
)

const (
	DayA = iota*12 + 13
	DayB
	DayC
)
