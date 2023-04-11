package main

/*Ejercicio 2 - Manipulando el body

Vamos a crear un endpoint llamado /saludo. Con una pequeña estructura con nombre y apellido que al pegarle deberá
responder en texto “Hola + nombre + apellido”

El endpoint deberá ser de método POST
Se deberá usar el package JSON para resolver el ejercicio
La respuesta deberá seguir esta estructura: “Hola Andrea Rivas”
La estructura deberá ser como esta:
{
		“nombre”: “Andrea”,
		“apellido”: “Rivas”
}
*/

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

func main() {

	// Creo router con gin
	router := gin.Default()

	// Captura la solicitud POST "/saludo" y devuelve "Hola + nombre + apellido"
	router.POST("/saludo", func(c *gin.Context) {

		var saludo User

		if err := c.ShouldBindJSON(&saludo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//c.String(200, "Hola "+j.Nombre+" "+j.Apellido)
		respuesta := "Hola " + saludo.Nombre + " " + saludo.Apellido
		c.JSON(http.StatusOK, gin.H{"message": respuesta})
	})

	// Inicia el servidor en el puerto 8080
	router.Run(":8080")

}
