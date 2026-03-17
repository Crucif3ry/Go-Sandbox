package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func hello(nom string) {
	fmt.Println("Salut ", nom)
}

func devinette() {

	tentative := 0
	devine := false
	random := rand.Intn(100)

	fmt.Println("Bienvenue dans GoDevinette")
	fmt.Println("Consigne : choisir un nombre entier entre 0 & 100")

	for !devine {

		tentative++

		var choix int

		fmt.Print("Entrer un nombre : ")
		fmt.Scanln(&choix)

		if choix == random {
			fmt.Println("Bravo c'est gagné en", tentative, "tentatives")
			devine = true
		} else if choix > random {
			fmt.Println("Non, c'est moins")
		} else if choix < random {
			fmt.Println("Non, c'est plus")
		}
	}
}

func request(ip string) {
	url := "https://api.country.is/" + ip + "?fields=city,postal,asn"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func task() {
	var notes = make(map[string]int)
	notes["Bob"] = 14
	notes["Alice"] = 19

	alive := true
	var choix int
	var eleve string
	var note int

	fmt.Println("Bienvenue dans Go-Notes un gestionnaire de notes")

	for alive {
		fmt.Println("1 = Voir les notes\n2 = Ajouter une note\n3 = Supprimer une note\n4 = Quitter le gestionnaire de note")
		fmt.Print("Entrer votre choix : ")
		fmt.Scanln(&choix)
		switch choix {
		case 1:
			fmt.Println("==================")
			for eleve := range notes {
				fmt.Println("La note de", eleve, "est", notes[eleve])
			}
			fmt.Println("==================")
		case 2:
			fmt.Print("Nom de l'élève : ")
			fmt.Scanln(&eleve)
			fmt.Print("Note entre 0 & 20 : ")
			fmt.Scanln(&note)
			if note > 20 {
				fmt.Println("La note doit etre entre 0 & 20")
				alive = false
			} else if note < 0 {
				fmt.Println("La note doit etre entre 0 & 20")
				alive = false
			}
			notes[eleve] = note
		case 3:
			fmt.Print("Eleve à supprimer : ")
			fmt.Scanln(&eleve)
			delete(notes, eleve)
		default:
			alive = false
		}
	}
}

// goroutines

var wg sync.WaitGroup // synchro les goroutines pour que le prog principale attande la fin des goroutines enfants

func run(name string) {
	defer wg.Done() // fin d'une goroutine ( s'execute a la fin de la fonction ) decremente de 1
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(name, " : ", time.Now())
	}
}

func execute() {
	debut := time.Now()

	wg.Add(1) // declare un go routine a attendre incremente de 1
	go run("Processus 1")
	wg.Add(1)
	go run("Processus 2")
	wg.Add(1)
	go run("Processus 3")

	wg.Wait() // attend que le nombre de goroutine soit a 0
	fin := time.Now()
	fmt.Println(fin.Sub(debut))
	fmt.Println("Comme on le voit on lance les 3 processus en parrallèle")
}

func main() {
	var choix int
	var prenom string
	var ip string
	alive := true

	for alive {
		fmt.Println("Bienvenue dans la Go-Sandbox \n Choix 1 = Jeu Hello \n Choix 2 = Jeu de devinette \n Choix 3 = Requete API sur une IP publique \n Choix 4 = Gestionnaire de notes \n Choix 5 = Processus en Goroutines \n Choix 6 = Quitter")
		fmt.Print("Entrer votre choix : ")
		fmt.Scanln(&choix)
		switch choix {
		case 1:
			fmt.Print("Entrer votre prénom : ")
			fmt.Scanln(&prenom)
			hello(prenom)
		case 2:
			devinette()
		case 3:
			fmt.Print("Entrer une ip publique (ex:77.1.2.3) : ")
			fmt.Scanln(&ip)
			request(ip)
		case 4:
			task()
		case 5:
			execute()
		default:
			alive = false
		}
	}
}
