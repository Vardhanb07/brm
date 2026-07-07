package brm

import (
	"testing"
)

func TestConvertToIntAndReset(t *testing.T) {
	tests := []struct {
		name string
		t    string
		n    int
		err  error
	}{
		{
			name: "convert '56' and reset to ''",
			t:    "56",
			n:    56,
			err:  nil,
		},
		{
			name: "fail to convert '56w'",
			t:    "56w",
			n:    0,
			err:  ErrParseDuration,
		},
	}
	t.Parallel()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			n, err := convertToIntAndReset(&test.t)
			if n != test.n && err != test.err {
				t.Errorf("expected: output: %v, err: %v, got: output: %v, err: %v", test.n, test.err, n, err)
			}
			if test.t != "" && err == nil {
				t.Errorf("expected: '', got: %v", test.t)
			}
		})
	}
}

func TestParseDuration(t *testing.T) {
	tests := []struct {
		name     string
		duration string
		mins     int
		hours    int
		days     int
		months   int
		err      error
	}{
		{
			name:     "convert '5min5d5h5m'",
			duration: "5min5d5h5m",
			mins:     5,
			hours:    5,
			days:     5,
			months:   5,
			err:      nil,
		},
		{
			name:     "fail to convert '5mins'",
			duration: "5mins",
			err:      ErrParseDuration,
		},
		{
			name:     "convert '5m5h5d5min'",
			duration: "5m5h5d5min",
			mins:     5,
			hours:    5,
			days:     5,
			months:   5,
			err:      nil,
		},
		{
			name:     "convert '5min5min'",
			duration: "5min5min",
			mins:     10,
			err:      nil,
		},
		{
			name:     "coonvert '5h5h'",
			duration: "5h5h",
			hours:    10,
			err:      nil,
		},
		{
			name:     "convert '5d5d'",
			duration: "5d5d",
			days:     10,
			err:      nil,
		},
		{
			name:     "convert '5m5m'",
			duration: "5m5m",
			months:   10,
			err:      nil,
		},
	}
	t.Parallel()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mins, hours, days, months, err := parseDuration(test.duration)
			if test.err == nil {
				if test.mins != mins {
					t.Errorf("expected: mins: %v, got: mins: %v", test.mins, mins)
				} else if test.hours != hours {
					t.Errorf("expected: hours: %v, got: hours: %v", test.hours, hours)
				} else if test.days != days {
					t.Errorf("expected: days: %v, got: days: %v", test.days, days)
				} else if test.months != months {
					t.Errorf("expected: months: %v, got: months: %v", test.months, months)
				}
			} else if err != test.err {
				t.Errorf("expected: err: %v, got: err: %v", test.err, err)
			}
		})
	}
}
