package hello

import "testing"

func TestSayHello(t *testing.T) {
	want := "Hello, mateusz"
	got := SayHello("mateusz")

	if want != got {
		t.Errorf("wanted %s but got %s", want, got)
	}
}
