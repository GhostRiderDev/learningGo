package test

import (
	"fmt"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	if os.Getenv("RUNTIME_ENV") == "development" {
		t.Skip("Skipping a test in development environment")
	}

	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	tests := []struct {
		a    int
		b    int
		want int
	}{
		{a: 1, b: 2, want: 3},
		{a: 7, b: 12, want: 19},
		{a: -3, b: 5, want: 2},
		{a: 0, b: 0, want: 0},
	}

	for _, tt := range tests {
		if got, want := AddInt(tt.a, tt.b), tt.want; got != want {
			t.Errorf("Invalid add result got %d\nExpected => %d", got, want)
		}
	}

	t.Run("Test Strings", func(t *testing.T) {
		t.Parallel()
		if got, want := AddString("Hello", "World"), "HelloWorld"; got != want {
			t.Errorf("Invalid add result got %s\nExpected => %s", got, want)
		}
	})
}

func AddInt(num1, num2 int) int {
	return num1 + num2
}

func AddString(str1, str2 string) string {
	return fmt.Sprintf("%s%s", str1, str2)
}
