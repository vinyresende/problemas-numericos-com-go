package main

import "fmt"

func main() {
	for n := 1; n <= 100; n++ {
		if n%3 == 0 {
			fmt.Print("Pin ")
		}

		if n%5 == 0 {
			fmt.Print("Pan ")
		}

		if n%3 != 0 && n%5 != 0 {
			fmt.Println(n)
		} else {
			fmt.Print("\n")
		}
	}
}
