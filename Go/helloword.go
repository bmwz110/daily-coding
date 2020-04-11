package main

import "fmt"

func variableZeroValue() {
	var a int
	var s string

	fmt.Printf("%q %q\n", a, s) 
}

func variableInitialValue() {
	var a, b = 3, "abc"

	fmt.Println(a, b)
}

func variableShorter() {
	a, b := 3, "def"
	b = "str"

	fmt.Println(a, b)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableShorter()
}

 
