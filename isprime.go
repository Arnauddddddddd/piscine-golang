package piscine

func IsPrime(nb int) bool {
	liste_prime := []int{2, 3, 5, 7, -2, -3, -5, -7}
	if nb < 0 {
		return false
	}
	if nb == 1 || nb == 2 || nb == 3 || nb == 5 || nb == 7 {
		return true
	}
	for i := 11; i < nb; i++ {
		if !(i%2 == 0 || i%3 == 0 || i%5 == 0 || i%7 == 0) {
			liste_prime = append(liste_prime, i)
		}
	}
	for j := range liste_prime {
		if nb%liste_prime[j] == 0 {
			return false
		}
	}
	return true
}
