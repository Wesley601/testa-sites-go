package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func sitesLog(site string, status bool) {
	logger(
		"log-status.txt",
		time.Now().Format("02/01/2006 15:04:05")+" - "+
			site+" - online "+strconv.FormatBool(status)+"\n",
	)
}

func errorsLog(err error) {
	logger("log-errors.txt", err.Error())
}

func logger(file string, message string) {
	arquivo, err := os.OpenFile("logs/"+file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer arquivo.Close()

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(message)
}
