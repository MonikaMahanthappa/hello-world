package helloworld_test

import (
	"github.com/stretchr/testify/assert"
	. "greeting/helloworld"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	t.Run("should greet with Hello, World!", func(t *testing.T) {
		expected := "Hello, World!"
		assert.Equal(t, expected, HelloWorld())
	})
}
