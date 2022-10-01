package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	//fmt.Println(quote.Glass()) //rsc.io/quote
	// Get a greeting message and print it.
	//message := greetings.Hello("Thud")
	//fmt.Println(message)

	//multipleinput
	//disHellos()
	//
	//disError()
	//array
	//disArray()
	//Loop
	//disForLoop()
	//userInput
	userInputScan()

	//Map
}

func disArray() {
	var arr [5]int                //declare value [0 0 0 0 0]
	arrB := [5]int{1, 5, 3, 4, 6} //declare init value
	arrC := []int{4, 3}           //declare slice array
	arrC = append(arrC, 7)        //append slice

	arr[4] = 10
	fmt.Println(arr) // [0 0 0 0 10]
	fmt.Println(arrB)
	fmt.Println(arrC) // [4 3 7]

	var strarr [3]string
	fmt.Println(strarr) // [  ]

	strarr[0], strarr[1], strarr[2] = "first", "second", "third"
	fmt.Println(strarr) //[first second third]
	//array tutorial
}
func disForLoop() {

	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println(index, " ", value)
	}
	for i := 0; i < 6; i++ {
		fmt.Println(i)
		if 0 == (i % 2) {
			//fmt.Println("even")
		} else if 0 != (i % 2) {
			//fmt.Println("odd")
		}
	}
	//loop
}
func disError() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello("thud")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
	//error handler tutorial
}
func disHellos() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	//slice of names
	names := []string{"Glarys", "Copper", "Ark"}

	messages, errs := greetings.Hellos(names)
	if errs != nil {
		log.Fatal(errs)
	}
	fmt.Println(messages)
}

func userInputScan() string {
	var input string
	fmt.Scan(&input)
	fmt.Println(input, &input)
	return input
}
