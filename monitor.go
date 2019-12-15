package main

import (
	"fmt"
)

func main() {

	fmt.Println("Who's monitoring?")
	var name string
	fmt.Scan(&name)

	fmt.Println("Hey", name, "choose an option")
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")

	var action int
	/* another option to read the input, + verbose and + ~complicated */
	fmt.Scanf("%d", &action)

	fmt.Println("Option chosen:", action)
}
