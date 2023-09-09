package piscine

func Rot14(s string) string {
	message_codee := ""
	for i := 0; i < len(s); i++ {
		if rune(s[i]) <= 90 && rune(s[i]) >= 65 {
			if rune(s[i])+14 >= 90 {
				message_codee += string(65 + rune(s[i]) + 14 - 90)
			} else {
				message_codee += string(rune(s[i])+14)
			}
		}
		if rune(s[i]) <= 122 && rune(s[i]) >= 97 {
			if rune(s[i])+14 >= 122 {
				message_codee += string(97 + rune(s[i]) + 14 - 123)
			} else {
				message_codee += string(rune(s[i]) + 14)
			}
		}
		if (rune(s[i]) > 122 || rune(s[i]) < 97) && (rune(s[i]) > 90 || rune(s[i]) < 65) {
			message_codee += string(s[i])
		}
	}
	return message_codee
}
