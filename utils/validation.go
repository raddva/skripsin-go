package utils

import (
	"fmt"
	. "skripsin/models"
)

// Validation Func
func IsNumber(s string) bool {
	if s == "" {
		return false
	}

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func IsAlpha(s string) bool {
	if s == "" {
		return false
	}

	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c < 'A' || c > 'Z') &&
			(c < 'a' || c > 'z') &&
			c != ' ' &&
			c != '.' &&
			c != ',' {
			return false
		}
	}
	return true
}

func IsValidYear(s string) bool {
	var year int
	if !IsNumber(s) {
		return false
	}

	fmt.Sscanf(s, "%d", &year)
	if year < 1900 || year > NYear { // batasi karena sekarang baru 2026 :v
		return false
	}
	return true
}

// Fungsi untuk mengecek apakah input adalah angka dan berada di range menu valid
func IsValidMenuChoice(input string, min int, max int) bool {
	if !IsNumber(input) {
		return false
	}

	var choice int
	fmt.Sscanf(input, "%d", &choice)
	if choice >= min && choice <= max {
		return true
	}
	return false
}

// End of validation func
