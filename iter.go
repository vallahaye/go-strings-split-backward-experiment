package strings

import (
	"iter"
	. "strings"
	"unicode/utf8"
)

// SplitBackwardSeq returns an iterator over all substrings of s separated by sep,
// starting from the end.
// The iterator yields the same strings that would be returned by [SplitBackward](s, sep),
// but without constructing the slice.
// It returns a single-use iterator.
func SplitBackwardSeq(s, sep string) iter.Seq[string] {
	return splitBackwardSeq(s, sep, len(sep))
}

// SplitBackwardAfterSeq returns an iterator over substrings of s split after each instance of sep,
// starting from the end.
// The iterator yields the same strings that would be returned by [SplitBackwardAfter](s, sep),
// but without constructing the slice.
// It returns a single-use iterator.
func SplitBackwardAfterSeq(s, sep string) iter.Seq[string] {
	return splitBackwardSeq(s, sep, 0)
}

// splitSeq is SplitBackwardSeq or SplitBackwardAfterSeq, configured by how
// many bytes of sep to include in the results (none or all).
func splitBackwardSeq(s, sep string, sepSave int) iter.Seq[string] {
	if len(sep) == 0 {
		return explodeBackwardSeq(s)
	}
	return func(yield func(string) bool) {
		for {
			i := LastIndex(s, sep)
			if i < 0 {
				break
			}
			frag := s[i+sepSave:]
			if !yield(frag) {
				return
			}
			s = s[:i]
		}
		yield(s)
	}
}

// explodeBackwardSeq returns an iterator over the runes in s, starting from
// the end.
func explodeBackwardSeq(s string) iter.Seq[string] {
	return func(yield func(string) bool) {
		for len(s) > 0 {
			_, size := utf8.DecodeLastRuneInString(s)
			i := len(s) - size
			if !yield(s[i:]) {
				return
			}
			s = s[:i]
		}
	}
}
