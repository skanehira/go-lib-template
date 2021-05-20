package lib

import "testing"

func TestHelloWorld(t *testing.T) {
	got := HelloWorld()
	want := "hello world"
	if want != got {
		t.Errorf("unexpected result. want: %s, got: %s", want, got)
	}
}
