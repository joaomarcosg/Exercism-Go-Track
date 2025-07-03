package greeting

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
	got := HelloWorld()

	if expected != got {
		t.Errorf("HelloWorld() = %q; want %q", got, expected)
	}
}
