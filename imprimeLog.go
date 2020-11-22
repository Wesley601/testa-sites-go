package main

import (
	"fmt"
	"io/ioutil"
)

func imprimeLog() {
	fmt.Println("Exibindo Logs...")
	arquivo, err := ioutil.ReadFile("logs/log-status.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
