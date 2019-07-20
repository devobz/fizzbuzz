package main

import "testing"

func TestFizzbuzz(t *testing.T) {
	t.Run("It should be 1", func(t *testing.T) {
		got := fizzbuzz(1)
		want := "1"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
	t.Run("It should be 2", func(t *testing.T) {
		got := fizzbuzz(2)
		want := "2"

		if got != want {
			t.Errorf("got '%s' want '%s',", got, want)
		}
	})
	t.Run("It should be Fizz", func(t *testing.T) {
		got := fizzbuzz(3)
		want := "Fizz"

		if got != want {
			t.Errorf("got '%s' want '%s',", got, want)
		}
	})
	t.Run("It should be 4", func(t *testing.T) {
		got := fizzbuzz(4)
		want := "4"

		if got != want {
			t.Errorf("got '%s' want '%s',", got, want)
		}
	})
	t.Run("It should be Buzz", func(t *testing.T) {
		got := fizzbuzz(5)
		want := "Buzz"

		if got != want {
			t.Errorf("got '%s' want '%s',", got, want)
		}
	})
	t.Run("It should be Fizz", func(t *testing.T) {
		got := fizzbuzz(6)
		want := "Fizz"

		if got != want {
			t.Errorf("got '%s' want '%s',", got, want)
		}
	})
	t.Run("It should be 7", func(t *testing.T) {
		got := fizzbuzz(7)
		want := "7"

		if got != want {
			t.Errorf("got '%s' want '%s',", got, want)
		}
	})
	t.Run("It should be 8", func(t *testing.T) {
		got := fizzbuzz(8)
		want := "8"

		if got != want {
			t.Errorf("got '%s' want '%s',", got, want)
		}
	})
	t.Run("It should be Fizz", func(t *testing.T) {
		got := fizzbuzz(9)
		want := "Fizz"

		if got != want {
			t.Errorf("got '%s' want '%s',", got, want)
		}
	})
	t.Run("It should be Buzz", func(t *testing.T) {
		got := fizzbuzz(10)
		want := "Buzz"

		if got != want {
			t.Errorf("got '%s' want '%s',", got, want)
		}
	})
	t.Run("It should be 11", func(t *testing.T) {
		got := fizzbuzz(11)
		want := "11"

		if got != want {
			t.Errorf("got '%s' want '%s',", got, want)
		}
	})
	t.Run("It should be Fizz", func(t *testing.T) {
		got := fizzbuzz(12)
		want := "Fizz"

		if got != want {
			t.Errorf("got '%s' want '%s',", got, want)
		}
	})
	t.Run("It should be 13", func(t *testing.T) {
		got := fizzbuzz(13)
		want := "13"

		if got != want {
			t.Errorf("got '%s' want '%s',", got, want)
		}
	})
	t.Run("It should be 14", func(t *testing.T) {
		got := fizzbuzz(14)
		want := "14"

		if got != want {
			t.Errorf("got '%s' want '%s',", got, want)
		}
	})
	t.Run("It should be FizzBuzz", func(t *testing.T) {
		got := fizzbuzz(15)
		want := "FizzBuzz"

		if got != want {
			t.Errorf("got '%s' want '%s',", got, want)
		}
	})
}
