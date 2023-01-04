package tool

import "testing"

func TestGreeting(t *testing.T) {
	got, want := Greeting("world"), "Hello world"
	if got != want {
		t.Errorf("Foo() = %v, want %v", got, want)
	}
}
