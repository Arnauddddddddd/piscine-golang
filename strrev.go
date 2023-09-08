package main

func StrRev(s string) string {
	listeinversee := ""
	for i := 0; i < len(s); i++ {
		listeinversee = string(s[i]) + listeinversee
	}
	return listeinversee
}
