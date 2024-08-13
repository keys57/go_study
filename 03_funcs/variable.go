package main

// 不定参数
func YourName(name string, aliases ...string) {
	//aliases是一个切片
}

func CallYourName() {
	YourName("小明")
	YourName("小明", "大明")
	YourName("小明", "大明", "肥明")
	aliases := []string{"大明", "肥明"}
	YourName("小明", aliases...)
}
