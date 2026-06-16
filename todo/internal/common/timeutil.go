package common

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/araddon/dateparse"
	"github.com/maniartech/gotime"
)

func ParseDateTime(input string) (time.Time, error) {
	if input == "" {
		return time.Time{}, fmt.Errorf("empty date/time")
	}
	input = strings.TrimSpace(input)
	if t, err := time.Parse(time.RFC3339, input); err == nil {
		return t, nil
	}
	if t, ok := parseRelativeDateTime(input, time.Now()); ok {
		return t, nil
	}
	t, err := dateparse.ParseLocal(input)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date/time %q: %w", input, err)
	}
	return t, nil
}

func ParseHumanDuration(input string) (time.Duration, error) {
	if strings.TrimSpace(input) == "" {
		return 0, fmt.Errorf("empty duration")
	}
	if d, err := time.ParseDuration(input); err == nil {
		return d, nil
	}
	value, unit, err := parseRelativeNumberUnit(input)
	if err != nil {
		return 0, err
	}
	base := time.Now()
	var target time.Time
	switch unit {
	case "s":
		target = base.Add(time.Duration(value) * time.Second)
	case "m":
		target = base.Add(time.Duration(value) * time.Minute)
	case "h":
		target = base.Add(time.Duration(value) * time.Hour)
	case "d":
		target = gotime.Days(value, base)
	case "w":
		target = gotime.Weeks(value, base)
	default:
		return 0, fmt.Errorf("unsupported duration unit in %q", input)
	}
	return target.Sub(base), nil
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

func parseRelativeDateTime(input string, base time.Time) (time.Time, bool) {
	v := strings.ToLower(strings.TrimSpace(input))
	switch v {
	case "now":
		return base, true
	case "today":
		return gotime.SoD(base), true
	case "tomorrow":
		return gotime.Tomorrow(), true
	case "yesterday":
		return gotime.Yesterday(), true
	case "next week":
		return gotime.NextWeek(), true
	case "last week":
		return gotime.LastWeek(), true
	case "next month":
		return gotime.NextMonth(), true
	case "last month":
		return gotime.LastMonth(), true
	case "next year":
		return gotime.NextYear(), true
	case "last year":
		return gotime.LastYear(), true
	}

	sign := 1
	if strings.Contains(v, " ago") {
		sign = -1
		v = strings.TrimSuffix(v, " ago")
	}
	v = strings.TrimPrefix(v, "in ")
	v = strings.TrimSuffix(v, " from now")

	n, unit, err := parseRelativeNumberUnit(v)
	if err != nil {
		return time.Time{}, false
	}
	n *= sign

	switch unit {
	case "s":
		return base.Add(time.Duration(n) * time.Second), true
	case "m":
		return base.Add(time.Duration(n) * time.Minute), true
	case "h":
		return base.Add(time.Duration(n) * time.Hour), true
	case "d":
		return gotime.Days(n, base), true
	case "w":
		return gotime.Weeks(n, base), true
	case "mo":
		return gotime.Months(n, base), true
	case "y":
		return gotime.Years(n, base), true
	default:
		return time.Time{}, false
	}
}

var relRe = regexp.MustCompile(`^([+-]?\d+)\s*([a-zA-Z]+)$`)

func parseRelativeNumberUnit(input string) (int, string, error) {
	compact := strings.ToLower(strings.TrimSpace(input))
	compact = strings.ReplaceAll(compact, " ", "")
	m := relRe.FindStringSubmatch(compact)
	if len(m) != 3 {
		return 0, "", fmt.Errorf("invalid relative value %q", input)
	}
	n, err := strconv.Atoi(m[1])
	if err != nil {
		return 0, "", fmt.Errorf("invalid numeric value in %q", input)
	}
	switch m[2] {
	case "s", "sec", "secs", "second", "seconds":
		return n, "s", nil
	case "m", "min", "mins", "minute", "minutes":
		return n, "m", nil
	case "h", "hr", "hrs", "hour", "hours":
		return n, "h", nil
	case "d", "day", "days":
		return n, "d", nil
	case "w", "wk", "wks", "week", "weeks":
		return n, "w", nil
	case "mo", "mon", "month", "months":
		return n, "mo", nil
	case "y", "yr", "yrs", "year", "years":
		return n, "y", nil
	default:
		return 0, "", fmt.Errorf("unsupported relative unit in %q", input)
	}
}
