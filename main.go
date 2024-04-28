package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// Endpoint para verificar si el servidor está funcionando
	router.GET("/alive", func(c *gin.Context) {
		c.String(http.StatusOK, "Alive")
	})

	// Un segundo endpoint para mostrar un mensaje específico
	router.GET("/obedece", func(c *gin.Context) {
		c.String(http.StatusOK, "sms:+50368304843|hola")
	})

	// Un endpoint raíz que simplemente devuelve una respuesta básica
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Root Endpoint")
	})

	// Inicia el servidor en el puerto 8080
	router.Run(":8080")
}
