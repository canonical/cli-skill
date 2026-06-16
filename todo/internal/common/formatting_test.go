package common

import "testing"

func TestDetectColorSupport(t *testing.T) {
	t.Run("respects no color", func(t *testing.T) {
		t.Setenv("NO_COLOR", "1")
		t.Setenv("TERM", "xterm-256color")
		t.Setenv("COLORTERM", "truecolor")

		c := DetectColorSupport()
		if c.Supported {
			t.Fatalf("expected color to be disabled when NO_COLOR is set")
		}
	})

	t.Run("supports 256 color terminals", func(t *testing.T) {
		t.Setenv("NO_COLOR", "")
		t.Setenv("TERM", "xterm-256color")
		t.Setenv("COLORTERM", "")

		c := DetectColorSupport()
		if !c.Supported {
			t.Fatalf("expected color support for xterm-256color")
		}
		if !c.Has256 {
			t.Fatalf("expected Has256 to be true")
		}
	})

	t.Run("supports truecolor terminals", func(t *testing.T) {
		t.Setenv("NO_COLOR", "")
		t.Setenv("TERM", "xterm")
		t.Setenv("COLORTERM", "truecolor")

		c := DetectColorSupport()
		if !c.Supported {
			t.Fatalf("expected color support for truecolor terminal")
		}
		if !c.HasTruecolor {
			t.Fatalf("expected HasTruecolor to be true")
		}
	})

	t.Run("dumb terminals are not supported", func(t *testing.T) {
		t.Setenv("NO_COLOR", "")
		t.Setenv("TERM", "dumb")
		t.Setenv("COLORTERM", "")

		c := DetectColorSupport()
		if c.Supported {
			t.Fatalf("expected dumb terminal to disable colors")
		}
	})
}

func TestFormattingHelpers(t *testing.T) {
	t.Run("color section uses color when supported", func(t *testing.T) {
		t.Setenv("NO_COLOR", "")
		t.Setenv("TERM", "xterm-256color")

		got := ColorSection("Usage:")
		want := "\033[1;30mUsage:\033[0m"
		if got != want {
			t.Fatalf("unexpected ColorSection output: got %q want %q", got, want)
		}
	})

	t.Run("color section falls back to bold for no color", func(t *testing.T) {
		t.Setenv("NO_COLOR", "1")
		t.Setenv("TERM", "xterm-256color")

		got := ColorSection("Usage:")
		want := "\033[1mUsage:\033[0m"
		if got != want {
			t.Fatalf("unexpected ColorSection fallback: got %q want %q", got, want)
		}
	})

	t.Run("format section mirrors color section behavior", func(t *testing.T) {
		t.Setenv("NO_COLOR", "")
		t.Setenv("TERM", "xterm-256color")

		got := FormatSection("Flags:")
		want := "\033[1;30mFlags:\033[0m"
		if got != want {
			t.Fatalf("unexpected FormatSection output: got %q want %q", got, want)
		}
	})

	t.Run("bold only when colors are disabled", func(t *testing.T) {
		t.Setenv("NO_COLOR", "1")
		t.Setenv("TERM", "xterm-256color")

		got := Bold("todo")
		want := "\033[1mtodo\033[0m"
		if got != want {
			t.Fatalf("unexpected Bold output: got %q want %q", got, want)
		}
	})

	t.Run("bold passes through with colors enabled", func(t *testing.T) {
		t.Setenv("NO_COLOR", "")
		t.Setenv("TERM", "xterm-256color")

		got := Bold("todo")
		if got != "todo" {
			t.Fatalf("expected plain text when colors are enabled, got %q", got)
		}
	})

	t.Run("strip formatting removes ansi escapes", func(t *testing.T) {
		in := "\033[1;30mUsage:\033[0m test \033[1mtext\033[0m"
		got := StripFormatting(in)
		if got != "Usage: test text" {
			t.Fatalf("unexpected StripFormatting result: got %q", got)
		}
	})
}
