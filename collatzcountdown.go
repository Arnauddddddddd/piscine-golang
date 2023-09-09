package piscine

func CollatzCountdown(start int) int {
	nombre := 0
	for i := 0; start != 1; {
		if start%2 == 0 {
			start = start / 2
		} else {
			start = start*3 + 1
		}
		i++
		nombre = i
	}
	return nombre
}
