package strings_test

import (
	"math"
	"math/rand"
	"slices"
	"testing"

	. "github.com/vallahaye/go-strings-split-backward-experiment"
)

var abcd = "abcd"
var faces = "☺☻☹"
var commas = "1,2,3,4"
var dots = "1....2....3....4"

type SplitTest struct {
	s   string
	sep string
	n   int
	a   []string
}

var splitbackwardtests = []SplitTest{
	{"", "", -1, []string{}},
	{abcd, "", 2, []string{"d", "abc"}},
	{abcd, "", 4, []string{"d", "c", "b", "a"}},
	{abcd, "", -1, []string{"d", "c", "b", "a"}},
	{faces, "", -1, []string{"☹", "☻", "☺"}},
	{faces, "", 3, []string{"☹", "☻", "☺"}},
	{faces, "", 17, []string{"☹", "☻", "☺"}},
	{"☺�☹", "", -1, []string{"☹", "�", "☺"}},
	{abcd, "a", 0, nil},
	{abcd, "a", -1, []string{"bcd", ""}},
	{abcd, "z", -1, []string{"abcd"}},
	{commas, ",", -1, []string{"4", "3", "2", "1"}},
	{dots, "...", -1, []string{"4", "3.", "2.", "1."}},
	{faces, "☹", -1, []string{"", "☺☻"}},
	{faces, "~", -1, []string{faces}},
	{"1 2 3 4", " ", 3, []string{"4", "3", "1 2"}},
	{"1 2", " ", 3, []string{"2", "1"}},
	{"", "T", math.MaxInt / 4, []string{""}},
	{"\xff-\xff", "", -1, []string{"\xff", "-", "\xff"}},
	{"\xff-\xff", "-", -1, []string{"\xff", "\xff"}},
}

func TestSplitBackward(t *testing.T) {
	for _, tt := range splitbackwardtests {
		a := SplitBackwardN(tt.s, tt.sep, tt.n)
		if !slices.Equal(a, tt.a) {
			t.Errorf("SplitBackward(%q, %q, %d) = %v; want %v", tt.s, tt.sep, tt.n, a, tt.a)
			continue
		}
		if tt.n < 0 {
			a2 := slices.Collect(SplitBackwardSeq(tt.s, tt.sep))
			if !slices.Equal(a2, tt.a) {
				t.Errorf(`collect(SplitBackwardSeq(%q, %q)) = %v; want %v`, tt.s, tt.sep, a2, tt.a)
			}
		}
		if tt.n == 0 {
			continue
		}
		// s := Join(a, tt.sep)
		// if s != tt.s {
		// 	t.Errorf("Join(SplitBackward(%q, %q, %d), %q) = %q", tt.s, tt.sep, tt.n, tt.sep, s)
		// }
		if tt.n < 0 {
			b := SplitBackward(tt.s, tt.sep)
			if !slices.Equal(a, b) {
				t.Errorf("SplitBackward disagrees with SplitBackwardN(%q, %q, %d) = %v; want %v", tt.s, tt.sep, tt.n, b, a)
			}
		}
	}
}

var splitbackwardaftertests = []SplitTest{
	{abcd, "a", -1, []string{"abcd", ""}},
	{abcd, "z", -1, []string{"abcd"}},
	{abcd, "", -1, []string{"d", "c", "b", "a"}},
	{commas, ",", -1, []string{",4", ",3", ",2", "1"}},
	{dots, "...", -1, []string{"...4", "...3.", "...2.", "1."}},
	{faces, "☹", -1, []string{"☹", "☺☻"}},
	{faces, "~", -1, []string{faces}},
	{faces, "", -1, []string{"☹", "☻", "☺"}},
	{"1 2 3 4", " ", 3, []string{" 4", " 3", "1 2"}},
	{"1 2 3", " ", 3, []string{" 3", " 2", "1"}},
	{"1 2", " ", 3, []string{" 2", "1"}},
	{"123", "", 2, []string{"3", "12"}},
	{"123", "", 17, []string{"3", "2", "1"}},
}

func TestSplitBackwardAfter(t *testing.T) {
	for _, tt := range splitbackwardaftertests {
		a := SplitBackwardAfterN(tt.s, tt.sep, tt.n)
		if !slices.Equal(a, tt.a) {
			t.Errorf(`SplitBackward(%q, %q, %d) = %v; want %v`, tt.s, tt.sep, tt.n, a, tt.a)
			continue
		}
		if tt.n < 0 {
			a2 := slices.Collect(SplitBackwardAfterSeq(tt.s, tt.sep))
			if !slices.Equal(a2, tt.a) {
				t.Errorf(`collect(SplitBackwardAfterSeq(%q, %q)) = %v; want %v`, tt.s, tt.sep, a2, tt.a)
			}
		}
		// s := Join(a, "")
		// if s != tt.s {
		// 	t.Errorf(`Join(SplitBackward(%q, %q, %d), %q) = %q`, tt.s, tt.sep, tt.n, tt.sep, s)
		// }
		if tt.n < 0 {
			b := SplitBackwardAfter(tt.s, tt.sep)
			if !slices.Equal(a, b) {
				t.Errorf("SplitBackwardAfter disagrees with SplitBackwardAfterN(%q, %q, %d) = %v; want %v", tt.s, tt.sep, tt.n, b, a)
			}
		}
	}
}

func makeBenchInputHard() string {
	tokens := [...]string{
		"<a>", "<p>", "<b>", "<strong>",
		"</a>", "</p>", "</b>", "</strong>",
		"hello", "world",
	}
	x := make([]byte, 0, 1<<20)
	for {
		i := rand.Intn(len(tokens))
		if len(x)+len(tokens[i]) >= 1<<20 {
			break
		}
		x = append(x, tokens[i]...)
	}
	return string(x)
}

var benchInputHard = makeBenchInputHard()

func BenchmarkSplitBackwardEmptySeparator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SplitBackward(benchInputHard, "")
	}
}

func BenchmarkSplitBackwardSingleByteSeparator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SplitBackward(benchInputHard, "/")
	}
}

func BenchmarkSplitBackwardMultiByteSeparator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SplitBackward(benchInputHard, "hello")
	}
}

func BenchmarkSplitBackwardNSingleByteSeparator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SplitBackwardN(benchInputHard, "/", 10)
	}
}

func BenchmarkSplitBackwardNMultiByteSeparator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SplitBackwardN(benchInputHard, "hello", 10)
	}
}
