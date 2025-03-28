package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Criar uma instância do router Gin com configuração padrão
	r := gin.Default()

	// Definir rota para o endpoint raiz "/"
	r.GET("/", func(c *gin.Context) {
		// Criar a resposta
		response := map[string]string{
			"message": "Hello Go",
		}

		// O Gin já configura automaticamente o Content-Type como application/json
		c.JSON(http.StatusOK, response)
	})

	// Iniciar o servidor na porta 8080
	log.Println("Servidor iniciado na porta 8080")
	log.Fatal(r.Run(":8080"))
}