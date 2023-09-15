package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this string"

	h := sha256.New()

	fmt.Println([]byte(s))
	fmt.Println(len(s), len([]byte(s)))
	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
