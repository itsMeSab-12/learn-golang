package iterations

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat a 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("abc", 10)
	fmt.Println(repeated)
	//Output:abcabcabcabcabcabcabcabcabcabc
}
