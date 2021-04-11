package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const PERSONES = 100
const SETMANES = 5

type NomPersona struct {
	Nom  string
	Sexe string
}

type Persones struct {
	Nom         string       `json:"nom"`
	Cognoms     string       `json:"cognoms"`
	Adreça      *Adreca      `json:"adreça"`
	Departament *Departament `json:"departament"`
	Setmanes    []Setmana    `json:"setmanes"`
}

type Adreca struct {
	Carrer   string `json:"carrer"`
	Numero   int    `json:"numero"`
	Poblacio string `json:"poblacio"`
}

type Departament struct {
	Nom    string `json:"nom"`
	Carrec string `json:"càrrec,omitempty"`
	Sou    int    `json:"sou"`
}

type Setmana struct {
	Numero       int      `json:"número"`
	Dies         []string `json:"dies"`
	Justificacio string   `json:"justificacio,omitempty"`
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

func generaDepartaments(departamentsPersona []string, carrecsPersona []string, carrecsAssignats map[string]bool) *Departament {

	departament := departamentsPersona[rand.Intn(len(departamentsPersona))]
	result := Departament{
		Nom:    departament,
		Carrec: "",
		Sou:    2000 + rand.Intn(5)*100,
	}

	_, ok := carrecsAssignats[departament]
	if !ok {
		result.Carrec = carrecsPersona[0]
		result.Sou += 1000 + rand.Intn(10)*100
		carrecsAssignats[departament] = true
	}

	return &result
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

	return nil, errors.New("ha sortir la poblacio impossible")
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func generaSetmanes(dies []string, excuses []string) []Setmana {
	setmanes := make([]Setmana, 0)

	rand.Shuffle(len(dies), func(i, j int) { dies[i], dies[j] = dies[j], dies[i] })

	diesTriats := dies[:2]

	seguentSetmana := 1

	for i := 1; i <= SETMANES; i++ {

		if seguentSetmana == i {

			diesPresencia := make([]string, len(diesTriats))
			copy(diesPresencia, diesTriats)
			excusa := ""
			seguentSetmana = i + 1

			// Comprova si s'ha escaquejat ...
			if rand.Intn(100) < 25 {
				diesPresencia = remove(diesPresencia, rand.Intn(len(diesPresencia)))
				excusaRand := excuses[rand.Intn(len(excuses))]
				excusaTmp := strings.Split(excusaRand, ",")

				excusa = excusaTmp[0]
				wait, _ := strconv.Atoi(excusaTmp[1])
				seguentSetmana += wait
			}

			setmana := Setmana{
				Numero:       i,
				Dies:         diesPresencia,
				Justificacio: excusa,
			}
			setmanes = append(setmanes, setmana)
		}
	}
	return setmanes
}

func main() {
	nomsPersona := readNoms("../data/noms.txt")
	cognomsPersona := readStringsFile("../data/cognoms.txt")
	departamentsPersona := readStringsFile("../data/departaments.txt")
	poblacions := readPoblacions("../data/poblacions.txt")
	carrecsPersona := readStringsFile("../data/carrecs.txt")

	excuses := readStringsFile("../data/excuses.txt")
	dies := readStringsFile("../data/dies.txt")

	carrecsAssignats := make(map[string]bool)

	rand.Seed(time.Now().UnixNano())

	persones := make([]Persones, PERSONES)

	for i := 0; i < PERSONES; i++ {
		nom, _ := generaNom(nomsPersona)
		cognoms := triaRandom(cognomsPersona) + " " + triaRandom(cognomsPersona)
		departament := generaDepartaments(departamentsPersona, carrecsPersona, carrecsAssignats)
		poblacio, _ := generaPoblacio(poblacions)
		setmanes := generaSetmanes(dies, excuses)

		persona := Persones{
			Nom:         nom,
			Cognoms:     cognoms,
			Departament: departament,
			Adreça:      poblacio,
			Setmanes:    setmanes,
		}

		persones[i] = persona

	}

	str, err := json.Marshal(persones)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(str))

}
