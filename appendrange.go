package piscine

func AppendRange(min, max int) []int {
	tab := []int{}
	if min >= max {
		tab = append(tab, nil)
		return tab
	}
	for i := min; i < max; i++ {
		tab = append(tab, i)
	}
	return tab
}
