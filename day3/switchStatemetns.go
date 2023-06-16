package main

import "fmt"

func main() {
	i := 2

	fmt.Println("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Printf("One")
	case 2:
		fmt.Printf("two")
	case 3:
		fmt.Printf("three")
	}
}
