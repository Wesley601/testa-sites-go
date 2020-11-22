package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

func leSiteArquivo() ([]string, error) {
	sites := make([]string, 0, 8)
	arquivo, err := os.Open("sites.txt")
	defer arquivo.Close()

	if err != nil {
		errorsLog(err)
		return sites, errors.New(
			"Verifique se o arquivo 'sites.txt' existe no diretório raiz",
		)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		if err == io.EOF {
			break
		}
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
	}

	return sites, nil
}
