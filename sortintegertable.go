package piscine

func SortIntegerTable(table []int) {
	newTable := []int{}
	for i := len(table)-1; i >= 0; i-- {
		newTable = append(newTable, table[i])
		table = table[:i]
	}
	for j := 0; j < len(newTable); j++ {
		table = append(table, newTable[j])
	}
}
