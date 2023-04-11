package main

/*Ejercicio 1 - Prueba de Ping
Vamos a crear una aplicación Web con el framework Gin que tenga un endpoint /ping que al pegarle responda un texto que diga “pong”
El endpoint deberá ser de método GET
La respuesta de “pong” deberá ser enviada como texto, NO como JSON
*/

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Creo router con gin
	router := gin.Default()

	// Captura la solicitud GET "/ping" y devuelve "pong"
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong") //Respuesta de texto
		//c.JSON(200, gin.H{"message": "pong"}) //Respuesta de JSON
	})

	// Inicia el servidor en el puerto 8080
	router.Run(":8080")

}
