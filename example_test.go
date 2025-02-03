package strings_test

import (
	"fmt"

	strings "github.com/vallahaye/go-strings-split-backward-experiment"
)

func ExampleSplitBackward() {
	fmt.Printf("%q\n", strings.SplitBackward("a,b,c", ","))
	fmt.Printf("%q\n", strings.SplitBackward("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.SplitBackward(" xyz ", ""))
	fmt.Printf("%q\n", strings.SplitBackward("", "Bernardo O'Higgins"))
	// Output:
	// ["c" "b" "a"]
	// ["canal panama" "plan " "man " ""]
	// [" " "z" "y" "x" " "]
	// [""]
}

func ExampleSplitBackwardN() {
	fmt.Printf("%q\n", strings.SplitBackwardN("a,b,c", ",", 2))
	z := strings.SplitBackwardN("a,b,c", ",", 0)
	fmt.Printf("%q (nil = %v)\n", z, z == nil)
	// Output:
	// ["c" "a,b"]
	// [] (nil = true)
}

func ExampleSplitBackwardAfter() {
	fmt.Printf("%q\n", strings.SplitBackwardAfter("a,b,c", ","))
	// Output: [",c" ",b" "a"]
}

func ExampleSplitBackwardAfterN() {
	fmt.Printf("%q\n", strings.SplitBackwardAfterN("a,b,c", ",", 2))
	// Output: [",c" "a,b"]
}
