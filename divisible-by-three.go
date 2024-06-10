package main

import "fmt"

func main() {
	fmt.Println("Múltiplos de 3:")

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	fmt.Println("Múltiplos de 3: Pin")
	fmt.Println("Múltiplos de 5: Pan")

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("Pin")
		} else if i%5 == 0 {
			fmt.Println("Pan")
		} else {
			fmt.Println(i)
		}
	}
}
