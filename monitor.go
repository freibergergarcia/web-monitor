package main

import (
	"fmt"
	"os"
)

func main() {

	whoAreYou()
	chooseAnOption()

	action := readOption()

	switch action {
	case 1:
		fmt.Println("Monitoring...")
	case 2:
		fmt.Println("Showing logs...")
	case 0:
		fmt.Println("Bye!")
		os.Exit(0)
	default:
		fmt.Println("Couldn't find the option", action)
		os.Exit(-1)
	}
}

func whoAreYou() {
	fmt.Println("Who's monitoring?")
	var name string
	fmt.Scan(&name)
	fmt.Println("Hey", name, "choose an option")
}

func chooseAnOption() {
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")
}

func readOption() int {
	var action int
	/* another option to read the input, + verbose and + ~complicated */
	fmt.Scanf("%d", &action)
	fmt.Println("Option chosen:", action)

	return action
}
