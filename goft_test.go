package goft

import "testing"

func TestHello(t *testing.T) {
	actual := Hello()
	expected := "Hello World"

	if actual != expected {
		t.Errorf("Unexpected response:\nGot: %s, Wanted: %s", actual, expected)
	}
}
