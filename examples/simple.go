package main

import "fmt"

const (
	A = iota
	B
	C
	D
	E
	F
)

const (
	M = iota
	N
	O
	P
	Q
	R
)

func main() {
	fmt.Println(B)
	fmt.Println(N)
}
