package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	mapaCursos := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	curso := ""
	for curso != "q" {
		var cargaHoraria int
		fmt.Print("Digite um curso ou digite 'q' para sair: ")
		scanner.Scan()
		curso = scanner.Text()
		if curso != "q" {
			fmt.Print("Digite a carga horário do curso: ")
			fmt.Scanf("%d \n", &cargaHoraria)
			mapaCursos[curso] = cargaHoraria
		}
	}
	fmt.Println("Cursos registrados: ")
	for nomeCurso, cargaHoraria := range mapaCursos {
		fmt.Printf(" - %s: %dh \n", nomeCurso, cargaHoraria)
	}
	curso = ""
	for curso != "q" {
		fmt.Print("Digite o nome do curso a ser excluído ou digite 'q' para cancelar: ")
		scanner.Scan()
		curso = scanner.Text()
		if curso != "q" {
			cargaHoraria, cursoExiste := mapaCursos[curso]
			if cursoExiste {
				delete(mapaCursos, curso)
				fmt.Printf("O curso %s com %dh foi removido \n", curso, cargaHoraria)
			} else {
				fmt.Println("O curso digitado não existe.")
			}
		}
	}
	fmt.Println("Cursos registrados: ")
	for nomeCurso, cargaHoraria := range mapaCursos {
		fmt.Printf(" - %s: %dh \n", nomeCurso, cargaHoraria)
	}
}
