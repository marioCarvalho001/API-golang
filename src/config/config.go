package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConecaoBanco = ""
	Porta              = 0
)

func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_port"))
	if erro != nil {
		Porta = 9000
	}

	StringConecaoBanco = fmt.Sprintf("%s:%s@tcp(database:3306)/%s",
		os.Getenv("DB_user"),
		os.Getenv("DB_senha"),
		os.Getenv("DB_nome"),
	)
}
