package greeting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
	assert.Equal(t, expected, HelloWorld(), "they should be equal")
}
