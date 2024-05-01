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
		//sms:+50368304843|hola
		c.String(http.StatusOK, "get:photos")
	})

	// Un endpoint raíz que simplemente devuelve una respuesta básica
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Root Endpoint")
	})

	// Nuevo endpoint para cargar archivos ZIP
	router.POST("/x", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo recibir el archivo"})
			return
		}

		// Limitar el tamaño del archivo a 20MB
		const maxFileSize = 20 << 20 // 20 MB
		if file.Size > maxFileSize {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El archivo excede el tamaño máximo permitido"})
			return
		}

		// Guardar el archivo en una ubicación segura
		dst := "./tmp/" + file.Filename
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el archivo"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Archivo cargado con éxito", "path": dst})
	})

	// Inicia el servidor en el puerto 8080
	router.Run(":8080")
}
