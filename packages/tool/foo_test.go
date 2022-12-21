package tool

import "testing"

func TestFoo(t *testing.T) {
	got, want := Foo(), "bar"
	if got != want {
		t.Errorf("Foo() = %v, want %v", got, want)
	}
}
