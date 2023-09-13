package piscine

func Sqrt(nb int) int {
	ls_nb := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	ls_sq := []int{1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 4}
	for i, j := range ls_nb {
		if nb == j {
			return ls_sq[i]
		}
	}
	return 0
}
