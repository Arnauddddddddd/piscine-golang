package piscine

func IsPrime(nb int) bool {
	liste_prime := []int{}
	for i := 0; i < nb+1; i++ {
		if !(i%2 == 0 || i%3 == 0 || i%5 == 0 || i%7 == 0) {
			liste_prime = append(liste_prime, i)
		}
	}
	for j := range liste_prime {
		if liste_prime[j] == nb {
			return true
		}
	}
	return false
}
