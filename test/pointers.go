package main

import "log"

func main() {
	var someString string = "some"
	log.Printf("The value of the string is: %s", someString)

	var longString string = someString + "String"
	log.Printf("The value of the string is: %s", someString)
	log.Printf("The value of the long string is: %s", longString)

	// To pass by reference, use the & (pointer)

	var anotherString string = someString + "String"
	log.Printf("The value of another string is: %s", anotherString)
}
