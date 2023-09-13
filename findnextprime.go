package piscine

func FindNextPrime(nb int) int {
	tab := []int{}
	if nb <= 1 {
		return 2
	}
	if nb >= 10000000 {
		for i := 2; i < nb/1000; i++ {
			if nb%i == 0 {
				tab = append(tab, i)
			}
		}
	} else {
		for i := 2; i < nb; i++ {
			if nb%i == 0 {
				tab = append(tab, i)
			}
		}
	}
	for i := 2; i < nb/2; i++ {
		if nb%i == 0 {
			tab = append(tab, i)
		}
	}
	if len(tab) == 0 {
		return nb
	}
	return FindNextPrime(nb + 1)
}
