package common

import (
	"os"
	"testing"

	"github.com/muesli/termenv"
)

func forceProfile(profile termenv.Profile) func() {
	return SetColorProfileResolverForTesting(func(*termenv.Output) termenv.Profile {
		return profile
	})
}

func TestDetectColorSupport(t *testing.T) {
	t.Run("respects no color", func(t *testing.T) {
		restore := forceProfile(termenv.ANSI256)
		defer restore()
		t.Setenv("NO_COLOR", "1")

		c := DetectColorSupport()
		if c.Supported {
			t.Fatalf("expected color to be disabled when NO_COLOR is set")
		}
	})

	t.Run("supports 256 color terminals", func(t *testing.T) {
		restore := forceProfile(termenv.ANSI256)
		defer restore()
		_ = os.Unsetenv("NO_COLOR")

		c := DetectColorSupport()
		if !c.Supported {
			t.Fatalf("expected color support for xterm-256color")
		}
		if !c.Has256 {
			t.Fatalf("expected Has256 to be true")
		}
	})

	t.Run("supports truecolor terminals", func(t *testing.T) {
		restore := forceProfile(termenv.TrueColor)
		defer restore()
		_ = os.Unsetenv("NO_COLOR")

		c := DetectColorSupport()
		if !c.Supported {
			t.Fatalf("expected color support for truecolor terminal")
		}
		if !c.HasTruecolor {
			t.Fatalf("expected HasTruecolor to be true")
		}
	})

	t.Run("dumb terminals are not supported", func(t *testing.T) {
		restore := forceProfile(termenv.Ascii)
		defer restore()
		_ = os.Unsetenv("NO_COLOR")
		t.Setenv("TERM", "dumb")

		c := DetectColorSupport()
		if c.Supported {
			t.Fatalf("expected dumb terminal to disable colors")
		}
	})
}

func TestFormattingHelpers(t *testing.T) {
	t.Run("color section uses bold when background is unknown", func(t *testing.T) {
		restoreProfile := forceProfile(termenv.ANSI256)
		defer restoreProfile()
		restoreBg := SetBackgroundModeDetectorForTesting(func(*termenv.Output) string { return "" })
		defer restoreBg()
		_ = os.Unsetenv("NO_COLOR")

		got := ColorSection("Usage:")
		want := "\033[1mUsage:\033[0m"
		if got != want {
			t.Fatalf("unexpected ColorSection output: got %q want %q", got, want)
		}
	})

	t.Run("color section uses light gray for dark backgrounds", func(t *testing.T) {
		restoreProfile := forceProfile(termenv.ANSI256)
		defer restoreProfile()
		restoreBg := SetBackgroundModeDetectorForTesting(func(*termenv.Output) string { return "dark" })
		defer restoreBg()
		_ = os.Unsetenv("NO_COLOR")

		got := ColorSection("Usage:")
		want := "\033[37mUsage:\033[0m"
		if got != want {
			t.Fatalf("unexpected ColorSection output for dark bg: got %q want %q", got, want)
		}
	})

	t.Run("color section uses dark gray for light backgrounds", func(t *testing.T) {
		restoreProfile := forceProfile(termenv.ANSI256)
		defer restoreProfile()
		restoreBg := SetBackgroundModeDetectorForTesting(func(*termenv.Output) string { return "light" })
		defer restoreBg()
		_ = os.Unsetenv("NO_COLOR")

		got := ColorSection("Usage:")
		want := "\033[30mUsage:\033[0m"
		if got != want {
			t.Fatalf("unexpected ColorSection output for light bg: got %q want %q", got, want)
		}
	})

	t.Run("color section falls back to bold for no color", func(t *testing.T) {
		restore := forceProfile(termenv.ANSI256)
		defer restore()
		t.Setenv("NO_COLOR", "1")

		got := ColorSection("Usage:")
		want := "Usage:"
		if got != want {
			t.Fatalf("unexpected ColorSection fallback: got %q want %q", got, want)
		}
	})

	t.Run("format section mirrors color section behavior", func(t *testing.T) {
		restoreProfile := forceProfile(termenv.ANSI256)
		defer restoreProfile()
		restoreBg := SetBackgroundModeDetectorForTesting(func(*termenv.Output) string { return "" })
		defer restoreBg()
		_ = os.Unsetenv("NO_COLOR")

		got := FormatSection("Flags:")
		want := "\033[1mFlags:\033[0m"
		if got != want {
			t.Fatalf("unexpected FormatSection output: got %q want %q", got, want)
		}
	})

	t.Run("bold only when colors are disabled", func(t *testing.T) {
		restore := forceProfile(termenv.ANSI256)
		defer restore()
		t.Setenv("NO_COLOR", "1")

		got := Bold("todo")
		want := "todo"
		if got != want {
			t.Fatalf("unexpected Bold output: got %q want %q", got, want)
		}
	})

	t.Run("bold uses ansi when colors enabled", func(t *testing.T) {
		restore := forceProfile(termenv.ANSI256)
		defer restore()
		_ = os.Unsetenv("NO_COLOR")

		got := Bold("todo")
		if got != "\033[1mtodo\033[0m" {
			t.Fatalf("expected bold ansi text when colors are enabled, got %q", got)
		}
	})

	t.Run("strip formatting removes ansi escapes", func(t *testing.T) {
		in := "\033[1;96mUsage:\033[0m test \033[1mtext\033[0m"
		got := StripFormatting(in)
		if got != "Usage: test text" {
			t.Fatalf("unexpected StripFormatting result: got %q", got)
		}
	})

	t.Run("render inline tags strips tags when colors disabled", func(t *testing.T) {
		restore := forceProfile(termenv.ANSI256)
		defer restore()
		t.Setenv("NO_COLOR", "1")

		got := RenderInlineTags("<b>todo</b> <green>ok</green>")
		if got != "todo ok" {
			t.Fatalf("unexpected plain output: got %q", got)
		}
	})

	t.Run("render inline tags uses ansi when colors enabled", func(t *testing.T) {
		restore := forceProfile(termenv.ANSI256)
		defer restore()
		_ = os.Unsetenv("NO_COLOR")

		got := RenderInlineTags("<green>ok</green>")
		if got != "\033[32mok\033[0m" {
			t.Fatalf("unexpected colored output: got %q", got)
		}
	})
}
