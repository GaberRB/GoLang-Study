// Dup2 exibe a contagem e o texto das linhas que aparecem mais de uma
// vez na entrada. ele lê de stdin ou de uma lista de arquivos nomeados.
// Exercicio exibir o nome do arquivo em que ocorre as linhas duplicadas
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	lineDuplicate := os.Args[1]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}

			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			fmt.Println("Exibindo o nome do arquivo: ", lineDuplicate)
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	//NOTA: ignorando erros em potencial de input.Err()
}
