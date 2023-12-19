package main // or whatever your package name is

import (
	"testing"
)

func TestFindStringInTag(t *testing.T) {
	tests := []struct {
		name      string
		html      string
		tag       string
		want      []string
		wantError bool
	}{
		{
			name:      "Single occurrence",
			html:      "<p>Hello World</p>",
			tag:       "p",
			want:      []string{"Hello World"},
			wantError: false,
		},
		{
			name:      "Multiple occurrences",
			html:      "<div>First</div><div>Second</div>",
			tag:       "div",
			want:      []string{"First", "Second"},
			wantError: false,
		},
		{
			name:      "No occurrences",
			html:      "<span>None</span>",
			tag:       "div",
			want:      nil,
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindStringInTag(tt.html, tt.tag)
			if (err != nil) != tt.wantError {
				t.Errorf("FindStringInTag() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if !compareStringSlices(got, tt.want) {
				t.Errorf("FindStringInTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

// compareStringSlices compares two slices of strings for equality.
func compareStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestFindContentByID(t *testing.T) {
	tests := []struct {
		name      string
		html      string
		id        string
		want      string
		wantError bool
	}{
		{
			name:      "Valid ID",
			html:      `<div id="test">Hello World</div>`,
			id:        "test",
			want:      "Hello World",
			wantError: false,
		},
		{
			name:      "ID Not Present",
			html:      `<div>Hello World</div>`,
			id:        "test",
			want:      "",
			wantError: false,
		},
		// Add more test cases as needed...
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindContentByID(tt.html, tt.id)
			if (err != nil) != tt.wantError {
				t.Errorf("FindContentByID() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if got != tt.want {
				t.Errorf("FindContentByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindContentByClass(t *testing.T) {
	tests := []struct {
		name      string
		html      string
		className string
		want      []string
		wantError bool
	}{
		{
			name:      "Single occurrence",
			html:      `<div class="test">Hello World</div>`,
			className: "test",
			want:      []string{"Hello World"},
			wantError: false,
		},
		{
			name:      "Multiple occurrences",
			html:      `<div class="test">First</div><div class="test">Second</div>`,
			className: "test",
			want:      []string{"First", "Second"},
			wantError: false,
		},
		{
			name:      "No occurrences",
			html:      `<div class="different">Hello World</div>`,
			className: "test",
			want:      nil,
			wantError: false,
		},
		// More test cases can be added here...
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindContentByClass(tt.html, tt.className)
			if (err != nil) != tt.wantError {
				t.Errorf("FindContentByClass() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if !equalStringSlices(got, tt.want) {
				t.Errorf("FindContentByClass() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// equalStringSlices checks if two slices of strings are equal.
func equalStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
