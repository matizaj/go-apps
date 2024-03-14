package hello

import "testing"

func TestSayHello(t *testing.T) {
	subtests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, world",
		},
		{
			items:  []string{"mateusz", "hania"},
			result: "Hello, mateusz, hania",
		},
	}

	for _, str := range subtests {
		if s := SayHello(str.items); s != str.result {
			t.Errorf("wanted %s (%v)  but got %s", str.result, str.items, s)
		}
	}

}
