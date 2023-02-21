//Echo2 exibe seus argumentos de linha de comando.
//modifique o programa para exibir o args por linha

package main

import (
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Println(arg)
	}

}
