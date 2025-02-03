# Backward string splitting experiment

This module hosts an experimental implementation of the `strings.SplitBackward` function and its derivative. The API is unstable and may change at any time. **Do not depend on this module.**

## Proposal

<!-- proposal: strings: add SplitBackward function and its derivative -->

```go
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
func SplitBackward(s, sep string) []string

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
func SplitBackwardN(s, sep string, n int) []string

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
func SplitBackwardAfter(s, sep string) []string

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
func SplitBackwardAfterN(s, sep string, n int) []string

// SplitBackwardSeq returns an iterator over all substrings of s separated by sep,
// starting from the end.
// The iterator yields the same strings that would be returned by [SplitBackward](s, sep),
// but without constructing the slice.
// It returns a single-use iterator.
func SplitBackwardSeq(s, sep string) iter.Seq[string]

// SplitBackwardAfterSeq returns an iterator over substrings of s split after each instance of sep,
// starting from the end.
// The iterator yields the same strings that would be returned by [SplitBackwardAfter](s, sep),
// but without constructing the slice.
// It returns a single-use iterator.
func SplitBackwardAfterSeq(s, sep string) iter.Seq[string]
```
