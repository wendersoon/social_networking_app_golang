package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
)

func main() {
	fmt.Println("WEbapp")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":8000", r))
}
