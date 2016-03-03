# dtcvt
golang data type convert

## useage
	
	package main

	import (
		"fmt"
		"github.com/jesusslim/dtcvt"
	)

	func main() {
		//to string
		a := []byte("sa")
		r := dtcvt.MustString(a, "gg")
		fmt.Println("To string:", r)
	
		//to int
		b := 123
		r2 := dtcvt.MustInt(b)
		fmt.Println("To int:", r2)
	
		//to float
		c := []byte("12.7")
		r3 := dtcvt.MustFloat64(c)
		fmt.Println("To float:", r3)
	}