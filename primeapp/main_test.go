package main

import (
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

	prompt()

	//close writer
	_ = w.Close()

	// reset os.stdout
	os.Stdout = oldOut

	// read output of prompt function from read pipe
	out, _ := io.ReadAll(r)

	// test output

	if strings.Trim(string(out), " ") != "->" {
		t.Errorf("incorrect prompt, expected -> got %s", string(out))
	}
}
