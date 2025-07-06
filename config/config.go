package config

import (
	"log"

	"github.com/joho/godotenv"
)

// CarregarConfiguracoes carrega as configurações necessárias para a aplicação
func CarregarConfiguracoes() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env", err)
	}

	ConectarBanco()

}
