package strfmt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func testDurationParser(t *testing.T, toParse string, expected time.Duration) {
	r, e := ParseDuration(toParse)
	assert.NoError(t, e)
	assert.Equal(t, expected, r)
}

func testDurationSQLScanner(t *testing.T, dur time.Duration) {
	values := []interface{}{int64(dur), float64(dur)}
	for _, value := range values {
		var result Duration
		err := result.Scan(value)
		assert.NoError(t, err)
		assert.Equal(t, dur, time.Duration(result))
	}
}

func TestDurationParser(t *testing.T) {
	testcases := map[string]time.Duration{

		// parse the short forms without spaces
		"1ns": 1 * time.Nanosecond,
		"1us": 1 * time.Microsecond,
		"1µs": 1 * time.Microsecond,
		"1ms": 1 * time.Millisecond,
		"1s":  1 * time.Second,
		"1m":  1 * time.Minute,
		"1h":  1 * time.Hour,
		"1hr": 1 * time.Hour,
		"1d":  24 * time.Hour,
		"1w":  7 * 24 * time.Hour,
		"1wk": 7 * 24 * time.Hour,

		// parse the long forms without spaces
		"1nanoseconds":  1 * time.Nanosecond,
		"1nanos":        1 * time.Nanosecond,
		"1microseconds": 1 * time.Microsecond,
		"1micros":       1 * time.Microsecond,
		"1millis":       1 * time.Millisecond,
		"1milliseconds": 1 * time.Millisecond,
		"1second":       1 * time.Second,
		"1sec":          1 * time.Second,
		"1min":          1 * time.Minute,
		"1minute":       1 * time.Minute,
		"1hour":         1 * time.Hour,
		"1day":          24 * time.Hour,
		"1week":         7 * 24 * time.Hour,

		// parse the short forms with spaces
		"1  ns": 1 * time.Nanosecond,
		"1  us": 1 * time.Microsecond,
		"1  µs": 1 * time.Microsecond,
		"1  ms": 1 * time.Millisecond,
		"1  s":  1 * time.Second,
		"1  m":  1 * time.Minute,
		"1  h":  1 * time.Hour,
		"1  hr": 1 * time.Hour,
		"1  d":  24 * time.Hour,
		"1  w":  7 * 24 * time.Hour,
		"1  wk": 7 * 24 * time.Hour,

		// parse the long forms without spaces
		"1  nanoseconds":  1 * time.Nanosecond,
		"1  nanos":        1 * time.Nanosecond,
		"1  microseconds": 1 * time.Microsecond,
		"1  micros":       1 * time.Microsecond,
		"1  millis":       1 * time.Millisecond,
		"1  milliseconds": 1 * time.Millisecond,
		"1  second":       1 * time.Second,
		"1  sec":          1 * time.Second,
		"1  min":          1 * time.Minute,
		"1  minute":       1 * time.Minute,
		"1  hour":         1 * time.Hour,
		"1  day":          24 * time.Hour,
		"1  week":         7 * 24 * time.Hour,
	}

	for str, dur := range testcases {
		testDurationParser(t, str, dur)
		testDurationSQLScanner(t, dur)
	}
}
