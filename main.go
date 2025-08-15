package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	// Carrega as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env")
		return
	}

	fmt.Println(os.Getenv("TESTE"))
}
