package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"Prime", 7, true, "7 is a prime number"},
		{"Not Prime", 8, false, "8 is not a prime number because it is divisible by 2"},
		{"Zero", 0, false, "0 is not a prime number"},
		{"One", 1, false, "1 is not a prime number"},
		{"Negative", -11, false, "negative numbers are not prime"},
	}

	for _, v := range primeTests {
		result, msg := isPrime(v.testNum)
		if v.expected && !result {
			t.Errorf("%s: expected result true got false", v.name)
		}
		if !v.expected && result {
			t.Errorf("%s: expected result true got false", v.name)
		}
		if msg != v.msg {
			t.Errorf("%s: expected '%s' but got '%s'", v.name, msg, v.msg)
		}
	}
}
func Test_prompt(t *testing.T) {
	// save a copy of os.stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	intro()

	//close writer
	_ = w.Close()

	// reset os.stdout
	os.Stdout = oldOut

	// read output of prompt function from read pipe
	out, _ := io.ReadAll(r)

	// test output

	if !strings.Contains(string(out), "Enter a whole number") {
		t.Errorf("incorrect info text, got %s", string(out))
	}
}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "Please enter a whole number!"},
		{name: "zero", input: "0", expected: "0 is not a prime number"},
		{name: "one", input: "1", expected: "1 is not a prime number"},
		{name: "two", input: "2", expected: "2 is a prime number"},
		{name: "three", input: "3", expected: "3 is a prime number"},
		{name: "negative", input: "-1", expected: "negative numbers are not prime"},
		{name: "typed", input: "three", expected: "Please enter a whole number!"},
		{name: "decimal", input: "1.1", expected: "Please enter a whole number!"},
		{name: "quit", input: "q", expected: ""},
		{name: "QUIT", input: "Q", expected: ""},
		{name: "greek", input: "alfávito", expected: "Please enter a whole number!"},
	}

	for _, e := range tests {
		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)
		if !strings.EqualFold(e.expected, res) {
			t.Errorf("%s: incorrect value expected: %s, got: %s", e.name, e.expected, res)
		}
	}
}

func Test_readUserInput(t *testing.T) {
	// channel
	doneChan := make(chan bool)
	// instance of io.reader
	var stdin bytes.Buffer

	// user pressing 1 and enter and q enter
	stdin.Write([]byte("1\nq\n"))
	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}
