package main

import "github.com/01-edu/z01"

func Rot14(s string) string {
	message_codee := ""

	for i := 0; i < len(s); i++ {
		message_codee += string(rune(s[i]+14))
	}
	return message_codee
}


func main() {
	result := Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

