package main

import (
	"fmt"
	"strconv"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var result []string

	for i := 100; i >= 1; i-- {
		if isPrime(i) {
			continue
		}

		if i%3 == 0 && i%5 == 0 {
			result = append(result, "FooBar")
		} else if i%3 == 0 {
			result = append(result, "Foo")
		} else if i%5 == 0 {
			result = append(result, "Bar")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}

	for _, num := range result {
		fmt.Print(num, " ")
	}
}
