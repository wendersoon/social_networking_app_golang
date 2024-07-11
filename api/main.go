package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Carrega variáveis de ambiente na aplicação
	config.Carregar()

	// Obtém as rotas da aplicação
	r := router.Gerar()

	//Configura a escuta da aplicação
	fmt.Printf("Escutando na porta %d\n...", config.Porta)
	log.Fatal(http.ListenAndServe(":5000", r))

}
