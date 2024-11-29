package main

import "fmt"

func main() {
	for n := 1; n <= 100; n++ {
		if n%3 == 0 {
			fmt.Println(n)
		}
	}
}
