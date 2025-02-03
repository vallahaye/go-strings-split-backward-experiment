package strings

import (
	. "strings"
	"unicode/utf8"
)

// SplitBackward slices s into all substrings separated by sep, starting from
// the end, and returns a slice of the substrings between those separators.
//
// If s does not contain sep and sep is not empty, SplitBackward returns a
// slice of length 1 whose only element is s.
//
// If sep is empty, SplitBackward splits after each UTF-8 sequence, starting
// from the end. If both s and sep are empty, SplitBackward returns an empty
// slice.
//
// It is equivalent to [SplitBackwardN] with a count of -1.
func SplitBackward(s, sep string) []string {
	return genSplitBackward(s, sep, len(sep), -1)
}

// SplitBackwardN slices s into substrings separated by sep, starting from the
// end, and returns a slice of the substrings between those separators.
//
// The count determines the number of substrings to return:
//   - n > 0: at most n substrings; the last substring will be the unsplit remainder;
//   - n == 0: the result is nil (zero substrings);
//   - n < 0: all substrings.
//
// Edge cases for s and sep (for example, empty strings) are handled
// as described in the documentation for [SplitBackward].
func SplitBackwardN(s, sep string, n int) []string {
	return genSplitBackward(s, sep, len(sep), n)
}

// SplitBackwardAfter slices s into all substrings after each instance of sep,
// starting from the end, and returns a slice of those substrings.
//
// If s does not contain sep and sep is not empty, SplitBackwardAfter returns
// a slice of length 1 whose only element is s.
//
// If sep is empty, SplitBackwardAfter splits after each UTF-8 sequence,
// starting from the end. If both s and sep are empty, SplitBackwardAfter
// returns an empty slice.
//
// It is equivalent to [SplitBackwardAfterN] with a count of -1.
func SplitBackwardAfter(s, sep string) []string {
	return genSplitBackward(s, sep, 0, -1)
}

// SplitBackwardAfterN slices s into substrings after each instance of sep,
// starting from the end, and returns a slice of those substrings.
//
// The count determines the number of substrings to return:
//   - n > 0: at most n substrings; the last substring will be the unsplit remainder;
//   - n == 0: the result is nil (zero substrings);
//   - n < 0: all substrings.
//
// Edge cases for s and sep (for example, empty strings) are handled
// as described in the documentation for [SplitBackwardAfter].
func SplitBackwardAfterN(s, sep string, n int) []string {
	return genSplitBackward(s, sep, 0, n)
}

// Generic backward split: splits after each instance of sep, starting from the
// end, including sepSave bytes of sep in the subarrays.
func genSplitBackward(s, sep string, sepSave, n int) []string {
	if n == 0 {
		return nil
	}
	if sep == "" {
		return explodeBackward(s, n)
	}
	if n < 0 {
		n = Count(s, sep) + 1
	}

	if n > len(s)+1 {
		n = len(s) + 1
	}
	a := make([]string, n)
	n--
	i := 0
	for i < n {
		m := LastIndex(s, sep)
		if m < 0 {
			break
		}
		a[i] = s[m+sepSave:]
		s = s[:m]
		i++
	}
	a[i] = s
	return a[:i+1]
}

// explodeBackward splits s into a slice of UTF-8 strings, starting from the
// end, one string per Unicode character up to a maximum of n (n < 0 means no
// limit).
// Invalid UTF-8 bytes are sliced individually.
func explodeBackward(s string, n int) []string {
	l := utf8.RuneCountInString(s)
	if n < 0 || n > l {
		n = l
	}
	a := make([]string, n)
	for i := 0; i < n-1; i++ {
		_, size := utf8.DecodeLastRuneInString(s)
		m := len(s) - size
		a[i] = s[m:]
		s = s[:m]
	}
	if n > 0 {
		a[n-1] = s
	}
	return a
}
