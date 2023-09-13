package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	for i := 0; i < n; i++ {
		for j:= 0; j < 10; j++ {
			z01.PrintRune(rune(j)+48)
		}
		z01.PrintRune(' ')
		
	}
	z01.PrintRune('\n')
}


func main() {
	PrintCombN(1)
	PrintCombN(3)
	PrintCombN(9)
}
