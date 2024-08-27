package main

import (
	"fmt"
)

func main() {

	// Prime Number Generator
	ch := make(chan int)
	Limit := 100 // Limiting the genrator to 10000 numbers
	go generateNumber(ch)
	for {
		Prime := <-ch
		if Prime > Limit {
			break
		}
		fmt.Println(Prime)
		ch1 := make(chan int)
		go filterPrime(ch, ch1, Prime)
		ch = ch1
	}

	// Parallel Sum

	a := 1
	b := 20

	ch2 := make(chan int, 10)
	middle := (a + b) / 2

	go parallelSum(a, middle, ch2)
	go parallelSum(middle+1, b, ch2)

	s1 := <-ch2
	s2 := <-ch2

	total := s1 + s2
	fmt.Println("Total sum is:", total)

}

func generateNumber(ch chan int) { // This function generates numbers from number 2
	for i := 2; ; i++ {
		ch <- i
	}
}

func filterPrime(input, output chan int, Prime int) { //Filter function to send the prime numbers to output channel
	for {
		i := <-input
		if i%Prime != 0 {
			output <- i
		}
	}
}

func parallelSum(a, b int, ch chan int) { // Parallel Sum function
	sum := 0
	for i := a; i <= b; i++ {
		sum = sum + i
	}
	ch <- sum
}
