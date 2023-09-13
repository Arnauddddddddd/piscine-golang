package piscine

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}
	liste_fiba := []int{0, 1}
	ind := 0
	for i := 0; i <= index; i++ {
		liste_fiba = append(liste_fiba, liste_fiba[len(liste_fiba)-2] + liste_fiba[len(liste_fiba)-1])
		ind = liste_fiba[i]
	}
	return ind
}
