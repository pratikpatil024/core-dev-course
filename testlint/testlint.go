package testlint

import "fmt"

func testlint() {
	fmt.Println("hello world")

	x := 0

	if x == 1 {
		x = 0
	}
	z := x + 2

	fmt.Println(z)
	y := "x"
	fmt.Println(y)
}

func Main() {
	testlint()
}