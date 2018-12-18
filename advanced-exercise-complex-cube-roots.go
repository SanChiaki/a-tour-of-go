package main

import (
	"fmt"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	z := complex128(1)
	diff := cmplx.Pow(z, 3) - x
	for cmplx.Abs(diff) > 1e-10 {
		z = z - (diff / (3 * cmplx.Pow(z, 2)))
		diff = cmplx.Pow(z, 3) - x
	}
	return z
}

func main() {
	fmt.Println(Cbrt(2))
}
