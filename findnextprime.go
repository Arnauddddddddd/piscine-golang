package piscine

func FindNextPrime(nb int) int {
	tab := 0
	if nb <= 1 {
		return 2
	}
	for i := 2; i < Sqrt(nb)+1; i++ {
		if nb%i == 0 {
			tab += 1
		}
	}
	if tab == 0 {
		return nb
	}
	return FindNextPrime(nb + 1)
}
