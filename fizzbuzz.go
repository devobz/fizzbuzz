package main

import (
	"fmt"
	"strconv"
)

func fizzorbuzz(i, divider int) bool {
	return i%divider == 0
}

// Run FizzBuzz
func fizzbuzz(i int) string {
	fizz := fizzorbuzz(i, 3)
	buzz := fizzorbuzz(i, 5)

	if fizz && buzz {
		return "FizzBuzz"
	}
	if fizz {
		return "Fizz"
	}
	if buzz {
		return "Buzz"
	}

	return strconv.Itoa(i)

	/**
	 * SWITCH CASE
	 */
	// fizz := fizzorbuzz(i, 3)
	// buzz := fizzorbuzz(i, 5)
	// switch {
	// case fizz && buzz:
	// 	return "FizzBuzz"
	// case fizz:
	// 	return "Fizz"
	// case buzz:
	// 	return "Buzz"
	// default:
	// 	return strconv.Itoa(i)
	// }

	/**
	 * IF/ELSE
	 */
	// if i%3 == 0 && i%5 == 0 {
	// 	return "FizzBuzz"
	// }
	// if i%3 == 0 {
	// 	return "Fizz"
	// }
	// if i%5 == 0 {
	// 	return "Buzz"
	// }
	// return strconv.Itoa(i)
}

func main() {
	fmt.Println(fizzbuzz(1))
}
