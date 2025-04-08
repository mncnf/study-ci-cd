package main

import "testing"

func TestGreet(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"Empty input", "", "Hello, World!"},
		{"With name", "John", "Hello, John!"},
		{"Another name", "Go", "Hello, Go!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := greet(tt.input)
			if got != tt.want {
				t.Errorf("greet(%q) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}
