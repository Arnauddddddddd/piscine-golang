package piscine

func ConcatParams(args []string) string {
	str := ""
	for i := range args {
		str += args[i]
		if !(i == len(args) - 1) {
			str += "\n"
		}
	}
	return str
}
