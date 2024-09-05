package main

import (
	"fmt"
	"os"
	//"github.com/01-edu"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("--insert\n    -i\n       This flag inserts the string into the string passed as argument.\n--order\n  -o\n        This flag will behave like a boolean, if it is called it will order the argument.")
	} else if args[0] == "-h" || args[0] == "--help" {
		fmt.Println("-insert\n -i\n       This flag inserts the string into the string passed as argument.\n--order\n  -o\n        This flag will behave like a boolean, if it is called it will order the argument.")
	}
	if len(args) > 0 {
		liste_commandes := []string{"--insert=", "--order", "-i", "-o"}
		indice := 0
		for i := 0; i < len(args); i++ {
			for j := 0; j < len(args[i]); j++ {
				indice = 0
				for m := 0; m < len(liste_commandes); m++ {
					for n := 0; n < len(liste_commandes[m]); n++ {
						fmt.Println(string(args[i][j]), string(liste_commandes[m][n]))
						if args[i][j] == liste_commandes[m][n] {
							indice += 1
							fmt.Println(string(args[i][j]), string(liste_commandes[m][n]))
						}
					}
					if indice == len(liste_commandes[m]) {
						fmt.Println("mdr ça marche, ou pas....")
					}
				}
			}
		}
	}

}