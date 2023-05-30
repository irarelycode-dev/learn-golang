package main

import (
	"fmt"
	"unicode/utf8"
)

func stringsandrunes() {
	const s = "vaathi coming"
	fmt.Println("len:", len(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x", s[i])
	}
	fmt.Println("Rune counting:", utf8.RuneCountInString(s))
	for i := 0; i < len(s); i++ {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Println("runeValue and width:", runeValue, width)
	}
}
