package strstrstr_test

import (
	"fmt"
	"testing"

	"github.com/kchugalinskiy/education2/strstrstr"
)

func ExampleSubstring() {
	// Output: world!
	fmt.Println(strstrstr.Substring("Hello, world!", "world"))
}

func TestSubstring(t *testing.T) {
	tests := []struct {
		s, sub, want string
	}{
		{"Hello, world!", "world", "world!"},
		{"", "", ""},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%q in %q", tt.sub, tt.s), func(t *testing.T) {
			got := strstrstr.Substring(tt.s, tt.sub)
			if got != tt.want {
				t.Errorf("Substring(%q, %q) = %q; want %q", tt.s, tt.sub, got, tt.want)
			}
		})
	}
}
