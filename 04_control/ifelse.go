package main

func IfOnly(age int) {
	if age >= 18 {
		println("已成年")
	}
}

func IfElse(age int) {
	if age >= 18 {
		println("已成年")
	} else {
		println("未成年")
	}
}

func IfElseIf(age int) {
	if age >= 18 {
		println("已成年")
	} else if age >= 6 {
		println("青年")
	} else {
		println("未成年")
	}
}

func IfNewVariable(start int, end int) string {
	if distance := end - start; distance > 100 {
		return "远了"
	} else if distance > 60 {
		return "还行"
	} else {
		return "ok"
	}
}
