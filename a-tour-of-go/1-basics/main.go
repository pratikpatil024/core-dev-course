// Ref: https://go.dev/tour/basics

package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

const Pi = 3.14
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

var c, python, java bool
var i1, j1 int = 1, 2
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	return
}

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	// Commenting as this is a weak random number generator and the Linter is giving error
	// Eoor: Use of weak random number generator (math/rand instead of crypto/rand)
	// fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	fmt.Println(math.Pi)

	fmt.Println(add(42, 13))
	fmt.Println(add2(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))

	var i int

	fmt.Println(i, c, python, java)

	var c1, python1, java1 = true, false, "no!"

	fmt.Println(i1, j1, c1, python1, java1)

	var i2, j2 int = 1, 2

	k2 := 3
	c2, python2, java2 := true, false, "no!"
	fmt.Println(i2, j2, k2, c2, python2, java2)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var i3 int

	var f3 float64

	var b3 bool

	var s3 string

	fmt.Printf("%v %v %v %q\n", i3, f3, b3, s3)

	var x4, y4 int = 3, 4

	var f4 float64 = math.Sqrt(float64(x4*x4 + y4*y4))

	var z4 uint = uint(f4)

	// var z4 uint = f4
	fmt.Println(x4, y4, z4)

	v1 := 42
	v2 := 3.142
	v3 := 0.867 + 0.5i

	fmt.Printf("v2 is of type %T\n", v2)
	fmt.Printf("v1 is of type %T\n", v1)
	fmt.Printf("v3 is of type %T\n", v3)

	const World = "世界"

	const Truth = true

	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	// fmt.Println(Small)
	// fmt.Println(Big)
	// fmt.Println(needInt(Big))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
