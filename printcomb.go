package main

import (
	"piscine"
	"github.com/01-edu/z01"
)

func PrintComb() {
	listeChiffre := [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	listeNombres := []string{}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				if i < j && j < k {
					listeNombres = append(listeNombres, (listeChiffre[i] + listeChiffre[j] + listeChiffre[k]))
				}
			}
		}
	}	
	for nombre := range listeNombres {
		print(listeNombres[nombre], " ")
	z01.PrintRune('\n')
	}
} 

func main() {
	piscine.PrintComb()
}
