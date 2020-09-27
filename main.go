package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

// Não é necessário produzir todos os 8 relatórios, apenas os relatórios
// “Número de vagas”, “Total de votos nominais” e mais 2 outros, à escolha;
// Os aparentemente mais faceis:
// • Candidatos eleitos (sempre indicado partido, número de votos e coligação, se houver)
// • Candidatos mais votados dentro do número de vagas;
// • Votos totalizados por coligação ou partido (quando um partido não estiver em coligação), número de
// candidatos eleitos;
// • Votos totalizados por partido, número de candidatos eleitos;

func main() {
	// 	in := `first_name,last_name,username
	// "Rob","Pike",rob
	// Ken,Thompson,ken
	// "Robert","Griesemer","gri"
	// `
	file, err := os.Open("divulga20.csv")
	check(err)

	reader := csv.NewReader(file)

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		check(err)

		fmt.Println(line[0])
	}

	defer file.Close()

	//r := csv.NewReader(strings.NewReader(in))
	// r := csv.NewReader(strings.NewReader(file))

	// for {
	// 	record, err := r.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Println(record)
	// }
}
