package main

import (
	"fmt"
	"math"
)

func main() {
	var a uint8 = math.MaxUint8
	fmt.Println(a)
	var b uint16 = math.MaxUint16
	fmt.Println(b)
	var c uint32 = math.MaxUint32
	fmt.Println(c)
	var d uint64 = math.MaxUint64
	fmt.Println(d)
	fmt.Println()

	var e uint8 = math.MaxInt8
	fmt.Println(e)
	var f uint16 = math.MaxInt16
	fmt.Println(f)
	var g uint32 = math.MaxInt32
	fmt.Println(g)
	var h uint64 = math.MaxInt64
	fmt.Println(h)
	fmt.Println()

	var i int = math.MaxInt
	fmt.Println(i)
	fmt.Println()

	// byte is alias for uint8.
	// Represents raw data.
	var myByte byte = 'A'
	fmt.Println(myByte)

	// rune is alias for int32.
	// Used for strings, they are unicode.
	var myRune rune = '😀'
	fmt.Println(myRune)
	fmt.Println()

	// Casting
	i = int(d)
	fmt.Printf("d = %v\nint(d) = %v\n", d, i)
}
