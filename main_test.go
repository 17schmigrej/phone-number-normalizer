package main

import "testing"

func TestNormalize(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{"1234567890", "1234567890"},
		{"123 926 7230", "1239267230"},
		{"(185) 936 7890", "1859367890"},
		{"(111) 456-9190", "1114569190"},
		{"123-456-7894", "1234567894"},
		{"123-456-7890", "1234567890"},
		{"1234567891", "1234567891"},
		{"(123)572-4261", "1235724261"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actual := normalize(tc.input)
			if actual != tc.output {
				t.Errorf("got %s; wanted %s", actual, tc.output)
			}
		})
	}
}
