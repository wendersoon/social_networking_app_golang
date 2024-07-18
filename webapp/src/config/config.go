package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (

	// APIURL é a url de comunicação com a API
	APIURL = ""

	// Porta onde a aplicação web escuta
	Porta = 0

	// HashKey é utilizada para autenticar o cookie
	HashKey []byte

	// BlockKey é usada para criptografar os dados do cookie
	BlockKey []byte
)

// Carregar inicializa as variáveis de ambiente
func Carregar() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Porta, err = strconv.Atoi(os.Getenv("PORTA"))
	if err != nil {
		log.Fatal(err)
	}

	APIURL = os.Getenv("API_URL")

	HashKey = []byte(os.Getenv("HASHKEY"))

	BlockKey = []byte(os.Getenv("BLOCKKEY"))

}
