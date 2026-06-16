package common

import (
	"os"
	"strings"
)

// TerminalColors represents color support capabilities
type TerminalColors struct {
	Supported    bool
	Has256       bool
	HasTruecolor bool
}

// DetectColorSupport checks if the terminal supports colors
// It checks:
// - NO_COLOR environment variable (disables colors if set)
// - TERM environment variable
// - Presence of COLORTERM
func DetectColorSupport() TerminalColors {
	// Respect NO_COLOR directive
	if os.Getenv("NO_COLOR") != "" {
		return TerminalColors{Supported: false}
	}

	term := os.Getenv("TERM")

	// Check for truecolor support
	colorterm := os.Getenv("COLORTERM")
	hasTruecolor := colorterm == "truecolor" || colorterm == "24bit"

	// Check for 256 color support
	has256 := strings.Contains(term, "256color")

	// Basic color support detection
	supported := term != "" && term != "dumb" && (has256 || hasTruecolor || strings.HasPrefix(term, "xterm") || strings.HasPrefix(term, "screen"))

	return TerminalColors{
		Supported:    supported,
		Has256:       has256,
		HasTruecolor: hasTruecolor,
	}
}

// FormatSection formats a section header with color or bold
// Uses ANSI escape codes for color if supported, otherwise uses bold
func FormatSection(text string) string {
	colors := DetectColorSupport()

	if colors.Supported {
		// Use bold dark gray for section headers
		// \033[1;30m = bold dark gray, \033[0m = reset
		return "\033[1;30m" + text + "\033[0m"
	}

	// Fallback to bold (without color)
	// \033[1m = bold, \033[0m = reset
	return "\033[1m" + text + "\033[0m"
}

// StripFormatting removes ANSI escape codes from a string
func StripFormatting(text string) string {
	// Simple regex to remove ANSI escape sequences
	result := ""
	i := 0
	for i < len(text) {
		if text[i] == '\033' {
			// Skip escape sequence
			i++
			// Find the end of the sequence
			if i < len(text) && text[i] == '[' {
				i++
				// Skip until we find a letter
				for i < len(text) && (text[i] >= '0' && text[i] <= '9' || text[i] == ';') {
					i++
				}
				if i < len(text) {
					i++ // skip the letter
				}
			}
		} else {
			result += string(text[i])
			i++
		}
	}
	return result
}

// Bold formats text with bold when colors aren't supported
func Bold(text string) string {
	colors := DetectColorSupport()
	if !colors.Supported {
		return "\033[1m" + text + "\033[0m"
	}
	return text
}

// ColorSection creates a formatted section header
func ColorSection(text string) string {
	colors := DetectColorSupport()
	if !colors.Supported {
		// Fallback to bold
		return "\033[1m" + text + "\033[0m"
	}
	// Use bold dark gray for section headers
	return "\033[1;30m" + text + "\033[0m"
}
