package piscine

func FindNextPrime(nb int) int {
	tab := []int{}
	if nb <= 1 {
		return 2
	}
	for i := 2; i < nb-1; i++ {
		if nb%i == 0 {
			tab = append(tab, i)
		}
	}
	if len(tab) == 0 {
		return nb
	}
	if nb <= 1000000000 {
		return FindNextPrime(nb + 1)
	}
	return 0
}
