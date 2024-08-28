package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	//print welcome
	intro()
	// create a channel to indicate user when user quit
	doneChan := make(chan bool)
	// start a goroutine to read user input and run program
	go readUserInput(os.Stdin, doneChan)
	// block until done channel gets a value
	<-doneChan
	// close the channel
	close(doneChan)
	// print goodbye
	fmt.Println("Goodbye")
}

func readUserInput(in io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(in)
	for {
		res, done := checkNumbers(scanner)
		if done {
			doneChan <- false
			return
		}
		fmt.Println(res)
		prompt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	scanner.Scan()
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a whole number!", false
	}

	_, msg := isPrime(num)
	return msg, false
}

func intro() {
	fmt.Println("Is it prime")
	fmt.Println("------------")
	fmt.Println("Enter a whole number to know if prime or not. Enter q to quit")
	prompt()
}

func prompt() {
	fmt.Println("->")
}

func isPrime(n int) (bool, string) {
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not a prime number", n)
	}

	if n < 0 {
		return false, "negative numbers are not prime"
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not a prime number because it is divisible by %d", n, i)
		}
	}
	return true, fmt.Sprintf("%d is a prime number", n)
}
