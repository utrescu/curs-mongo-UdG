package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const PERSONES = 100

type NomPersona struct {
	Nom  string
	Sexe string
}

type Assistencia struct {
	Nom         string
	Cognom      string
	Sexe        string
	Departament string
	Idiomes     []string
}

type Persones struct {
	Nom         string   `json:"nom"`
	Sexe        string   `json:"sexe,omitempty"`
	Adreça      *Adreca  `json:"adreça"`
	Departament string   `json:"departament"`
	Idiomes     []string `json:"idiomes"`
	Delegat     string   `json:"delegatSindical,omitempty"`
}

type Adreca struct {
	Carrer   string `json:"carrer"`
	Numero   int    `json:"numero"`
	Poblacio string `json:"poblacio"`
}

func readNoms(fileName string) []NomPersona {
	var noms []NomPersona

	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("error opening %s: %v", fileName, err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for i := 0; scanner.Scan(); i++ {
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error reading from file:", err)
			os.Exit(3)
		}
		linia := strings.Split(scanner.Text(), ",")
		noms = append(noms, NomPersona{Nom: linia[0], Sexe: linia[1]})

	}
	return noms
}

func readStringsFile(fileName string) []string {
	var noms []string

	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("error opening %s: %v", fileName, err))
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for i := 0; scanner.Scan(); i++ {
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error reading from file:", err)
			os.Exit(3)
		}
		linia := scanner.Text()
		noms = append(noms, linia)

	}
	return noms
}

func readPoblacions(fileName string) map[string]string {
	poblacions := make(map[string]string)

	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("error opening %s: %v", fileName, err))
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for i := 0; scanner.Scan(); i++ {
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error reading from file:", err)
			os.Exit(3)
		}
		linia := strings.Split(scanner.Text(), ", ")
		poble := linia[0]
		for _, carrer := range linia[1:] {
			poblacions[carrer] = poble
		}

	}
	return poblacions
}

func generaNom(nomsPersona []NomPersona) (string, string) {
	numNoms := len(nomsPersona)

	nom := nomsPersona[rand.Intn(numNoms)]

	return nom.Nom, nom.Sexe
}

func triaRandom(llista []string) string {
	triat := llista[rand.Intn(len(llista))]
	return triat
}

func generaIdiomes(idiomesPersona []string) []string {
	idiomes := make([]string, len(idiomesPersona))

	copy(idiomes, idiomesPersona)
	numIdiomes := rand.Intn(3) + 1

	rand.Shuffle(len(idiomes), func(i, j int) { idiomes[i], idiomes[j] = idiomes[j], idiomes[i] })

	parlaCatala := rand.Intn(100)
	if parlaCatala > 5 {
		idiomes = append([]string{"català"}, idiomes...)
	}
	if parlaCatala == 6 {
		idiomes = append([]string{"klingon"}, idiomes...)
		numIdiomes++
	}
	return idiomes[:numIdiomes]
}

func generaPoblacio(poblacions map[string]string) (*Adreca, error) {
	numero := rand.Intn(PERSONES/10) + 1
	element := rand.Intn(len(poblacions))

	i := 0
	for key, value := range poblacions {
		if element == i {
			return &Adreca{Carrer: key, Poblacio: value, Numero: numero}, nil
		}
		i++
	}

	return nil, errors.New("Ha sortir la poblacio impossible")
}

func calculaDelegats() int {
	switch {
	case PERSONES > 250:
		return 13
	case PERSONES > 100:
		return 9
	case PERSONES > 49:
		return 5
	case PERSONES > 30:
		return 3
	default:
		return 1
	}
}

func main() {
	nomsPersona := readNoms("noms.txt")
	cognomsPersona := readStringsFile("cognoms.txt")
	departamentsPersona := readStringsFile("departaments.txt")
	idiomesPersona := readStringsFile("idiomes.txt")
	poblacions := readPoblacions("poblacions.txt")

	delegatsSindicals := calculaDelegats()
	delegatsSindicalsAssignats := 0

	rand.Seed(time.Now().UnixNano())

	persones := make([]Persones, PERSONES)

	for i := 0; i < PERSONES; i++ {
		nom, sexe := generaNom(nomsPersona)
		cognom := triaRandom(cognomsPersona)
		departament := triaRandom(departamentsPersona)
		idiomes := generaIdiomes(idiomesPersona)
		poblacio, _ := generaPoblacio(poblacions)
		delegat := ""

		if delegatsSindicalsAssignats < delegatsSindicals {
			esDelegat := rand.Intn(10) < 1
			if esDelegat {
				delegat = "Si"
				delegatsSindicalsAssignats++
			}
		}

		persona := Persones{
			Nom:         fmt.Sprintf("%s %s", nom, cognom),
			Sexe:        sexe,
			Departament: departament,
			Idiomes:     idiomes,
			Adreça:      poblacio,
			Delegat:     delegat,
		}

		persones[i] = persona

	}

	str, err := json.Marshal(persones)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(str))

}
