package main

import (
	"testing"
	"unicode/utf8"

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

// go test -fuzz=Fuzz / -fuzztime
func FuzzReverse(f *testing.F) {
	testcase := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcase {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev := stringutil.Reverse(orig)
		doubleRev := stringutil.Reverse(rev)
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})

}
