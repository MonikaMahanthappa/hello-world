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

}
