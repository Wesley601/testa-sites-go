package main

import (
	"fmt"
	"time"
)

const monitoramento = 3
const deplay = 2

func monitorar() {
	sites, err := leSiteArquivo()

	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i <= monitoramento; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		time.Sleep(deplay * time.Second)
		fmt.Println("")
	}
}
