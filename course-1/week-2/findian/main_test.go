package main

import (
	"testing"
)

// TestFoundTrue tests strings that should return true
func TestFoundTrue(t *testing.T) {
	strings := [...]string{"ian", "Ian", "iuiygaygn", "I d skd a efju N"}

	for _, str := range strings {
		if found(str) == false {
			t.Errorf("Found was incorrect, got: false, want: true.")
		}
	}
}

// TestFoundFalse tests strings that should return false
func TestFoundFalse(t *testing.T) {
	strings := [...]string{"ihhhhhn", "ina", "xian"}

	for _, str := range strings {
		if found(str) == true {
			t.Errorf("Found was incorrect, got: true, want: false.")
		}
	}
}
