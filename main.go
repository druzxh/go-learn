package main

import "fmt"

func main() {
	//init var
	var name string

	// print text
	fmt.Println("Hello world")

	// assign var
	name = "Badru"

	// print var
	fmt.Println(name)

	// variable
	iniVariabel := "Isinya variable"
	fmt.Println(iniVariabel)

	// double variable
	var (
		firstName = "Jack"
		lastName  = "Kie"
	)

	// Join var
	fullName := firstName + lastName

	fmt.Println(fullName)

	// number

	fmt.Println("satu", 1)
	fmt.Println("dua", 2)
	fmt.Println("float", 2.5)

	// boolean
	fmt.Println("true", true)
	fmt.Println("false", false)

	// string
	fmt.Println("Ini adalah string")

	// func string
	fmt.Println("panjang", len("Panjang string"))
	fmt.Println("string ke 5"[1])

	// constant
	const dataName = "Data constant"

	// multiple constant
	const (
		firstConst  = "Data first"
		secondConst = "Data second"
	)

	// convert data

	var value32 int = 32768
	var value64 int64 = int64(value32)
	var value16 int16 = int16(value32)

	fmt.Println("val32", value32)
	fmt.Println("val64", value64)
	fmt.Println("val16", value16)

}
