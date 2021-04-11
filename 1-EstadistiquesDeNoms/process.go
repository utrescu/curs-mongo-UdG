package main

// TODO: Separar noms que són en el mateix registre: català i castellà?

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type comarca struct {
	Posicio     int
	Comarca     string  `json:"comarca"`
	Quantitat   int     `json:"quantitat"`
	Percentatge float64 `json:"percentatge"`
}

type persona struct {
	Nom       string    `json:"nom"`
	Sexe      string    `json:"sexe"`
	Total     int       `json:"total"`
	Comarques []comarca `json:"comarques"`
}

func main() {

	fitxers := []string{"Alt Empordà", "Baix Empordà", "Cerdanya", "Garrotxa", "Gironès", "Pla de l'Estany", "Ripollès", "Selva"}

	noms := make(map[string]persona)

	// Processar els CSV
	for _, fitxer := range fitxers {

		csvfile, err := os.Open(fitxer + ".csv")
		if err != nil {
			log.Fatalln("Couldn't open the csv file", err)
		}
		defer csvfile.Close()

		r := csv.NewReader(csvfile)
		r.Comma = ';'
		r.Comment = '#'

		for {
			// Read each record from csv
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			posicio, err := strconv.Atoi(record[0])
			if err != nil {
				log.Fatal("Si no és un número no és una posició " + record[0])
			}
			nomPersona := record[2]
			sexe := record[1]
			quantitat, err := strconv.Atoi(record[3])
			if err != nil {
				log.Fatal("Un número que no és un número!")
			}
			percentagestring := strings.Replace(record[4], ",", ".", 1)
			percentatge, err := strconv.ParseFloat(percentagestring, 64)
			if err != nil {
				log.Fatal("Percentatge erròni :" + percentagestring)
			}

			var clau = sexe + "-" + nomPersona

			novaComarca := comarca{Posicio: posicio, Comarca: fitxer, Quantitat: quantitat, Percentatge: percentatge}

			valor, ok := noms[clau]
			if !ok {
				noms[clau] = persona{Nom: nomPersona, Sexe: sexe, Total: quantitat, Comarques: []comarca{novaComarca}}
			} else {
				valor.Comarques = append(valor.Comarques, novaComarca)
				valor.Total += quantitat
				noms[clau] = valor
			}

		}

	}

	// Obtenir l'array
	values := make([]persona, 0, len(noms))
	for _, v := range noms {
		values = append(values, v)
	}

	// Marshall
	// jsonInfo, _ := json.Marshal(values)
	jsonInfo, _ := json.MarshalIndent(values, "", "   ")
	fmt.Printf("%s", jsonInfo)

}
