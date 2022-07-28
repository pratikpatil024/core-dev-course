package main

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p3 = &Vertex{1, 2} // has type *Vertex
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

type Vertex1 struct {
	Lat, Long float64
}

var m map[string]Vertex1

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice1(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dy)
	for i, _ := range res {
		res[i] = make([]uint8, dx)

		for j, _ := range res[i] {
			res[i][j] = uint8((i + j) / 2)
		}
	}
	return res
}

func WordCount(s string) map[string]int {
	res := make(map[string]int, 0)
	split := strings.Fields(s)
	for _, word := range split {
		count, ok := res[word]
		if ok {
			res[word] = count + 1
		} else {
			res[word] = 1
		}
	}
	return res
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	i, j := 0, 1
	return func() int {
		i, j = j, i+j
		return i
	}
}

func main() {
	i, j := 42, 2701

	p := &i // point to i
	fmt.Println(p)
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	v1 := Vertex{1, 2}
	p1 := &v1
	p1.X = 1e9
	fmt.Println(v1)

	fmt.Println(v1, p3, v2, v3)

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a1 := names[0:2]
	b1 := names[1:3]
	fmt.Println(a1, b1)

	b1[0] = "XXX"
	fmt.Println(a1, b1)
	fmt.Println(names)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s1 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s1)

	s2 := []int{2, 3, 5, 7, 11, 13}

	s2 = s2[1:4]
	fmt.Println(s2)

	s2 = s2[:2]
	fmt.Println(s2)

	s2 = s2[1:]
	fmt.Println(s2)

	s3 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s3)

	// Slice the slice to give it zero length.
	s3 = s3[:0]
	printSlice(s3)

	// Extend its length.
	s3 = s3[:4]
	printSlice(s3)

	// Drop its first two values.
	s3 = s3[2:]
	printSlice(s3)

	var s4 []int
	fmt.Println(s4, len(s4), cap(s4))
	if s4 == nil {
		fmt.Println("nil!")
	}

	a2 := make([]int, 5)
	printSlice1("a2", a2)

	b2 := make([]int, 0, 5)
	printSlice1("b2", b2)

	c2 := b2[:2]
	printSlice1("c2", c2)

	d2 := c2[2:5]
	printSlice1("d2", d2)

	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	var s5 []int
	printSlice(s5)

	// append works on nil slices.
	s5 = append(s5, 0)
	printSlice(s5)

	// The slice grows as needed.
	s5 = append(s5, 1)
	printSlice(s5)

	// We can add more than one element at a time.
	s5 = append(s5, 2, 3, 4)
	printSlice(s5)

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow1 := make([]int, 10)
	for i := range pow1 {
		pow1[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow1 {
		fmt.Printf("%d\n", value)
	}

	pic.Show(Pic)

	m = make(map[string]Vertex1)
	m["Bell Labs"] = Vertex1{
		40.68433, -74.39967,
	}
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])

	var m1 = map[string]Vertex1{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
	fmt.Println(m1)

	m3 := make(map[string]int)

	m3["Answer"] = 42
	fmt.Println("The value:", m3["Answer"])

	m3["Answer"] = 48
	fmt.Println("The value:", m3["Answer"])

	delete(m3, "Answer")
	fmt.Println("The value:", m3["Answer"])

	val, ok := m3["Answer"]
	fmt.Println("The value:", val, "Present?", ok)

	wc.Test(WordCount)

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
