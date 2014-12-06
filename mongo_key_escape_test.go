package mongo

import "testing"
import "github.com/bmizerany/assert"

func TestEscape(t *testing.T) {
	assert.Equal(t, "hello\uFF04world", Escape("hello$world"))
	assert.Equal(t, "hello\uFF0Eworld", Escape("hello.world"))
}

func TestUnescape(t *testing.T) {
	assert.Equal(t, "hello$world", Unescape("hello\uFF04world"))
	assert.Equal(t, "hello.world", Unescape("hello\uFF0Eworld"))
}
