package piscine

func AppendRange(min, max int) []int {
	tab := []int{}
	if min >= max {
		return tab
	}
	for i := min; i < max; i++ {
		tab = append(tab, i)
	}
	return tab
}
