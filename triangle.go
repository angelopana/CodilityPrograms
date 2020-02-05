package main

import "fmt"

// Simple little half triangle to full triangle nest stuff

func triangle() {

	for i := 0; i < 5; i++ {
		for j := 5; j > i; j-- {
			fmt.Printf(" ")
		}
		for k := 0; k < i; k++ {
			fmt.Printf("* ")
		}
		fmt.Println()
	}
} // end of main

func halfTriangle() {
	for i := 0; i < 5; i++ {
		for k := 0; k < i; k++ {
			fmt.Printf("* ")
		}
		fmt.Println()
	}
}
