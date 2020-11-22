package main

import (
	"fmt"
	"net/http"
)

func testaSite(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Erro no site", site, "Err", err)
		return
	}

	if response.StatusCode == 200 {
		fmt.Println("Site,", site, "carregado com sucesso!")
	} else {
		fmt.Println("Algo não deu muito certo, status code:", response.StatusCode)
	}
	sitesLog(site, response.StatusCode == 200)
}
