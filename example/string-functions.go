package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {

	p("Contains:  ", s.Contains("testing", "es"))
	p("Count:     ", s.Count("testing", "t"))
	p("HasPrefix: ", s.HasPrefix("testing", "te"))
	p("HasSuffix: ", s.HasSuffix("testing", "st"))
	p("Index:     ", s.Index("testing", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("testing"))
	p()

	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}
