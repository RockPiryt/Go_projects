package main

import "testing"

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative numbers", -13, false, "Negative numbers are not prime, by definition"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		print(result)

		// expected true for 7
		if e.expected && !result {
			t.Errorf("%s: expected true  and result false wtedy error", e.name)
		}

		// expected false for 8
		// && tylko do wartości bool
		if !e.expected && result {
			t.Errorf("%s: expected false but got true in test wtedy error", e.name)
		}

		if e.msg != msg {
			t.Errorf("expected %s but got in test %s", e.msg, msg)
		}

	}
}
