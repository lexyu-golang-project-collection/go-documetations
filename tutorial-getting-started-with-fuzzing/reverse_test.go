package main

import (
	"testing"

	"golang.org/x/example/stringutil"
)

func TestReverse(t *testing.T) {
	testcase := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}

	for _, tc := range testcase {
		rev := stringutil.Reverse(tc.in)
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}

}
