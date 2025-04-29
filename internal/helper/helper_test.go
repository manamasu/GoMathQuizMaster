package helper

import (
	"testing"
	"time"
)

// Its not testing what we "see" but rather that it doesn't crash
func TestTypewriterEffect(t *testing.T) {
	TypewriterEffect(50*time.Millisecond, "Hello, World!")
}
