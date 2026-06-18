package common

import (
	"testing"
	"time"
)

func TestParseDateTime(t *testing.T) {
	t.Run("parses rfc3339", func(t *testing.T) {
		in := "2026-06-16T10:00:00Z"
		got, err := ParseDateTime(in)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got.Format(time.RFC3339) != in {
			t.Fatalf("unexpected parse result: got %s want %s", got.Format(time.RFC3339), in)
		}
	})

	t.Run("parses tomorrow", func(t *testing.T) {
		got, err := ParseDateTime("tomorrow")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		now := time.Now()
		if !got.After(now) {
			t.Fatalf("expected tomorrow to be in the future, got %v", got)
		}
	})

	t.Run("rejects empty input", func(t *testing.T) {
		if _, err := ParseDateTime("   "); err == nil {
			t.Fatalf("expected error for empty input")
		}
	})
}

func TestParseHumanDuration(t *testing.T) {
	tests := []struct {
		name  string
		in    string
		min   time.Duration
		max   time.Duration
		isErr bool
	}{
		{name: "std duration", in: "2h", min: 2 * time.Hour, max: 2*time.Hour + 2*time.Second},
		{name: "days shorthand", in: "1d", min: 23 * time.Hour, max: 25 * time.Hour},
		{name: "weeks shorthand", in: "2w", min: 13 * 24 * time.Hour, max: 15 * 24 * time.Hour},
			{name: "months shorthand currently unsupported", in: "1mo", isErr: true},
			{name: "years shorthand currently unsupported", in: "1y", isErr: true},
		{name: "invalid", in: "banana", isErr: true},
		{name: "empty", in: "   ", isErr: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseHumanDuration(tc.in)
			if tc.isErr {
				if err == nil {
					t.Fatalf("expected error for %q", tc.in)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error for %q: %v", tc.in, err)
			}
			if got < tc.min || got > tc.max {
				t.Fatalf("duration out of expected range for %q: got %v expected [%v..%v]", tc.in, got, tc.min, tc.max)
			}
		})
	}
}
