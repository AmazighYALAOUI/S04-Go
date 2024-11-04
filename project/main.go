package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var choix1, choix2 string

	fmt.Println("Voulez-vous créer une tâche ? [Y/N]")
	fmt.Scan(&choix1)
	if choix1 == "n" {
		fmt.Println("Merci et à bientôt sur Task !")
	} else {
		fmt.Println("Quelle tâche souhaitez-vous ajouter ? [Ajouter/Modifier/Supprimer]")
		fmt.Scan(&choix2)
		if choix2 == "ajouter" {
			fmt.Println("Entrer le nom de votre tâche: ")
			//fmt.Scan(&Task)
			scanner := bufio.NewScanner(os.Stdin) // Créer un scanner pour lire l'entrée standard
			scanner.Scan()                        // Lire la ligne entière
			Task := scanner.Text()                // Récupérer la ligne lue
			fmt.Println("La liste de vôs tâches:", Task)
		}
	}

}
