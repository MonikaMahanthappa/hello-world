package fizzbuzz_test

import (
	"github.com/stretchr/testify/assert"
	"greeting/fizzbuzz"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("should be 1 when number is 1", func(t *testing.T) {
		assert.Equal(t, "1", fizzbuzz.FizzBuzz(1))
	})

	t.Run("should be 4 when number is 4", func(t *testing.T) {
		assert.Equal(t, "4", fizzbuzz.FizzBuzz(4))
	})

	t.Run("multiple of three", func(t *testing.T) {
		t.Run("should be Fizz when number is 3", func(t *testing.T) {
			assert.Equal(t, "Fizz", fizzbuzz.FizzBuzz(3))
		})

		t.Run("should be Fizz when number is 6", func(t *testing.T) {
			assert.Equal(t, "Fizz", fizzbuzz.FizzBuzz(6))
		})
	})

	t.Run("multiple of five", func(t *testing.T) {
		t.Run("should be Buzz when number is 5", func(t *testing.T) {
			assert.Equal(t, "Buzz", fizzbuzz.FizzBuzz(5))
		})
		t.Run("should be Buzz when number is 10", func(t *testing.T) {
			assert.Equal(t, "Buzz", fizzbuzz.FizzBuzz(10))
		})
	})

}
