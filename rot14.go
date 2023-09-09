package piscine

func Rot14(s string) string {
	message_codee := ""

	for i := 0; i < len(s); i++ {
		message_codee += string(rune(s[i] + 14))
	}
	return message_codee
}
