package atexit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanup(t *testing.T) {
	assert := assert.New(t)
	exit := new(AtExit)

	a := 0
	b := 0
	exit.Add(func() {
		a = 4
	})
	exit.Add(func() {
		b = 3
	})
	exit.Cleanup()

	assert.Equal(4, a)
	assert.Equal(3, b)
}
