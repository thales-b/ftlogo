package mytypes

import "strings"

type MyInt int

type MyString string

type MyBuilder struct {
	Contents strings.Builder
	Len      int
}

type StringUppercaser struct {
	Contents strings.Builder
}

// Twice multiplies its receiver by 2 and returns
// the result.
func (i MyInt) Twice() MyInt {
	return i * 2
}

// MyStringLen returns the length of s
func (s MyString) MyStringLen() int {
	return len(s)
}

func (mb MyBuilder) Hello() string {
	return "Hello, Gophers!"
}

func (su StringUppercaser) ToUpper() string {
	return strings.ToUpper(su.Contents.String())
}

func (i *MyInt) Double() {
	*i *= 2
}
