package common

import (
	"fmt"
	"time"

	"github.com/araddon/dateparse"
)

func ParseDateTime(input string) (time.Time, error) {
	if input == "" {
		return time.Time{}, fmt.Errorf("empty date/time")
	}
	if t, err := time.Parse(time.RFC3339, input); err == nil {
		return t, nil
	}
	t, err := dateparse.ParseLocal(input)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date/time %q: %w", input, err)
	}
	return t, nil
}

func FormatTime(t *time.Time, useRFC3339 bool) string {
	if t == nil {
		return ""
	}
	if useRFC3339 {
		return t.Format(time.RFC3339)
	}
	return t.Local().Format("2006-01-02 15:04 MST")
}
