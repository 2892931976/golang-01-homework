package main

func main() {
	var b bool
	// b = 0|1,编译是不通过的
	b = true
	b = false
	b = ("hello" == "world")
	if b {

	}

}
