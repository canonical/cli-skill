package common

import (
	"os"
	"strconv"
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

// detectBackground detects if the terminal has a dark or light background.
// Returns "dark", "light", or "" if unknown.
func detectBackground() string {
	// Check COLORFGBG environment variable (set by many terminal emulators)
	// Format is "foreground;background" where background is the ANSI color index
	if colorfgbg := os.Getenv("COLORFGBG"); colorfgbg != "" {
		parts := strings.Split(colorfgbg, ";")
		if len(parts) >= 2 {
			if bg, err := strconv.Atoi(parts[len(parts)-1]); err == nil {
				// 0 = black, 8 = dark gray → dark background
				// 7 = white, 15 = bright white → light background
				if bg < 8 {
					return "dark"
				}
				return "light"
			}
		}
	}
	return ""
}

// FormatSection formats a section header with color or bold
// Uses ANSI escape codes for color if supported, otherwise uses bold
func FormatSection(text string) string {
	// FormatSection uses the same logic as ColorSection
	return ColorSection(text)
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

// ColorSection creates a formatted section header.
// Uses light gray on dark backgrounds, dark gray on light backgrounds,
// and bold with the terminal's default foreground when background is unknown.
func ColorSection(text string) string {
	colors := DetectColorSupport()
	if !colors.Supported {
		return "\033[1m" + text + "\033[0m"
	}

	switch detectBackground() {
	case "dark":
		// Light gray foreground — readable on dark backgrounds
		return "\033[37m" + text + "\033[0m"
	case "light":
		// Dark foreground — readable on light backgrounds
		return "\033[30m" + text + "\033[0m"
	default:
		// Unknown background: use bold with terminal's default foreground color
		// so it reads well on both dark and light backgrounds
		return "\033[1m" + text + "\033[0m"
	}
}
