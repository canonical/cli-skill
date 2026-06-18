package common

import (
	"os"
	"strings"

	"github.com/muesli/termenv"
)

// TerminalColors represents color support capabilities
type TerminalColors struct {
	Supported    bool
	Has256       bool
	HasTruecolor bool
}

var newTermOutput = func() *termenv.Output {
	return termenv.NewOutput(os.Stdout)
}

var resolveColorProfile = func(out *termenv.Output) termenv.Profile {
	return out.ColorProfile()
}

var detectBackgroundMode = func(out *termenv.Output) string {
	if out.HasDarkBackground() {
		return "dark"
	}
	if out.ColorProfile() != termenv.Ascii {
		return "light"
	}
	return ""
}

// SetTermOutputFactoryForTesting overrides terminal output detection for tests.
// It returns a restore function that must be called to reset the default behavior.
func SetTermOutputFactoryForTesting(factory func() *termenv.Output) func() {
	prev := newTermOutput
	newTermOutput = factory
	return func() {
		newTermOutput = prev
	}
}

// SetColorProfileResolverForTesting overrides profile resolution for tests.
// It returns a restore function that must be called to reset the default behavior.
func SetColorProfileResolverForTesting(resolver func(*termenv.Output) termenv.Profile) func() {
	prev := resolveColorProfile
	resolveColorProfile = resolver
	return func() {
		resolveColorProfile = prev
	}
}

// SetBackgroundModeDetectorForTesting overrides background mode detection for tests.
// It returns a restore function that must be called to reset the default behavior.
func SetBackgroundModeDetectorForTesting(detector func(*termenv.Output) string) func() {
	prev := detectBackgroundMode
	detectBackgroundMode = detector
	return func() {
		detectBackgroundMode = prev
	}
}

var tagToANSI = map[string]string{
	"<b>":      "\033[1m",
	"</b>":     "\033[0m",
	"<light>":  "\033[37m",
	"</light>": "\033[0m",
	"<dark>":   "\033[30m",
	"</dark>":  "\033[0m",
	"<green>":  "\033[32m",
	"</green>": "\033[0m",
}

var knownTags = []string{
	"<b>", "</b>",
	"<light>", "</light>",
	"<dark>", "</dark>",
	"<green>", "</green>",
}

// DetectColorSupport checks terminal color capabilities via termenv.
func DetectColorSupport() TerminalColors {
	out := newTermOutput()
	if out.EnvNoColor() {
		return TerminalColors{Supported: false}
	}

	profile := resolveColorProfile(out)
	supported := profile != termenv.Ascii
	hasTruecolor := profile == termenv.TrueColor
	has256 := profile == termenv.ANSI256 || hasTruecolor

	return TerminalColors{
		Supported:    supported,
		Has256:       has256,
		HasTruecolor: hasTruecolor,
	}
}

// StripInlineTags removes known style tags such as <b>, <light>, and <green>.
func StripInlineTags(text string) string {
	out := text
	for _, tag := range knownTags {
		out = strings.ReplaceAll(out, tag, "")
	}
	return out
}

// RenderInlineTags converts inline style tags into ANSI escapes when supported.
// If color output is not supported, tags are removed and plain text is returned.
func RenderInlineTags(text string) string {
	if !DetectColorSupport().Supported {
		return StripInlineTags(text)
	}

out := text
	for tag, ansi := range tagToANSI {
		out = strings.ReplaceAll(out, tag, ansi)
	}
	return out
}

// detectBackground uses termenv to detect dark/light terminal background.
func detectBackground() string {
	if !DetectColorSupport().Supported {
		return ""
	}

	out := newTermOutput()
	if mode := detectBackgroundMode(out); mode != "" {
		return mode
	}

	if resolveColorProfile(out) == termenv.Ascii {
		fallbackOut := termenv.NewOutput(os.Stdout, termenv.WithTTY(true))
		if mode := detectBackgroundMode(fallbackOut); mode != "" {
			return mode
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
	return StripInlineTags(result)
}

// Bold formats text with bold when colors aren't supported
func Bold(text string) string {
	return RenderInlineTags("<b>" + text + "</b>")
}

// ColorSection creates a formatted section header.
// Uses light gray on dark backgrounds, dark gray on light backgrounds,
// and bold with the terminal's default foreground when background is unknown.
func ColorSection(text string) string {
	colors := DetectColorSupport()
	if !colors.Supported {
		return text
	}

	switch detectBackground() {
	case "dark":
		return RenderInlineTags("<light>" + text + "</light>")
	case "light":
		return RenderInlineTags("<dark>" + text + "</dark>")
	default:
		return RenderInlineTags("<b>" + text + "</b>")
	}
}
