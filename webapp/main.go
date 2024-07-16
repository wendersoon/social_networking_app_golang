package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {

	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Println("Escutando na porta 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
