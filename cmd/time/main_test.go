package main

import "testing"

func TestTZOf(t *testing.T) {
	dat := []struct {
		i  string
		o  string
		tf bool
	}{
		{"Singapore", "Asia/Singapore", true},
		{"sgp", "Asia/Singapore", true},
		{"what is the time in SGP?", "Asia/Singapore", true},
		{"nonsense location", "UTC", false},
		{"los Angeles", "America/Los_Angeles", true},
		{"New York", "America/New_York", true},
	}

	for i, d := range dat {
		tz, supported := tzOf(d.i)
		if supported != d.tf {
			t.Errorf("case: %d: expected response to be: %v, got: %v", i, d.tf, supported)
		}

		if tz != d.o {
			t.Errorf("case: %d: expected: %q, got: %q", i, d.o, tz)
		}
	}
}
