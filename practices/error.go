package main

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error is ", err)
		os.Exit(-1)
	}
}

func foo() {
	/*err := doStuff1()
	checkError(err)

	err = doStuff2()
	checkError(err)

	err = doStuff3()
	checkError(err)*/
}

func handleA() {
	fmt.Println("handle A")
}
func handleB() {
	fmt.Println("handle B")
}

func foo1() {
	var err error
	defer func() {
		if err != nil {
			handleA()
			handleB()
		}
	}()

	/*err = doStuff1()
	if err != nil {
		return
	}

	err = doStuff2()
	if err != nil {
		return
	}

	err = doStuff3()
	if err != nil {
		return
	}*/
}