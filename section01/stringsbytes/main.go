package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// contains()
	// checkStringEquality()
	// checkByteEquality()
	// compareBytes()
	// trimSpaces()
	// trimFunc()
	// replacer()
	// join()
	reader()
}

func contains() {
	fmt.Printf("%v\n", strings.Contains("hello world", "gophers"))
	fmt.Printf("%v\n", bytes.Contains([]byte("abc"), []byte("b")))
}

func indexFunc() {

}

func checkStringEquality() {
	fmt.Printf("%v\n", "a" == "b")
	fmt.Printf("%v\n", strings.ToUpper("a") == "A")
	fmt.Printf("%v\n", strings.ToUpper("ϵ") == "Ε")
}
func checkByteEquality() {
	fmt.Printf("%v\n", bytes.Equal([]byte{'a', 'b'}, []byte("ab")))
}

func checkCaseInsensitiveEquality() {

}

func compareBytes() {
	fmt.Printf("%v\n", bytes.Compare([]byte{'a', 'b'}, []byte("ab")))
	fmt.Printf("%v\n", bytes.Compare([]byte{'a', 'b'}, []byte("abc")))
	fmt.Printf("%v\n", bytes.Compare([]byte{'a', 'b', 'c'}, []byte("ab")))
}

func trimSpaces() {
	fmt.Printf("%q\n", strings.TrimSpace(" hello "))
	fmt.Printf("%v\n", bytes.TrimSpace([]byte(" hello ")))
	fmt.Printf("%q\n", bytes.TrimSpace([]byte(" hello ")))
}

func trimFunc() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	fmt.Printf("[%q]\n", strings.TrimFunc("    1234gophers unite5678", f))
}

func replacer() {
	r := strings.NewReplacer("alpha", "Α", "theta", "Θ", "delta", "Δ")
	fmt.Printf("%q\n", r.Replace("The alpha differs from the theta which differs from the delta"))
}

func join() {
	alphabet := " alpha beta   gamma "
	fields := strings.Fields(alphabet)
	fmt.Printf("Fields: %q\n", fields)
	fmt.Printf("Joined: %q\n", strings.Join(fields, "|"))
}

func reader() {
	var a, b, c string
	s := "a b c"
	r := strings.NewReader(s)
	_, err := fmt.Fscan(r, &a, &b, &c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("a: %q\nb: %q\nc: %q\n", a, b, c)
}
