package main

import (
	"github.com/gin-gonic/gin"
)

//Definimos una pseudobase de datos donde consultaremos la información

var empleados = map[string]string{
	"644": "Juan",
	"755": "Pedro",
	"777": "Luis",
}

func main() {
	server := gin.Default()

	server.GET("/", PaginaPrincipal)
	server.GET("/empleados/:id", BuscarEmpleado)
	server.Run(":8080")

}

//Handler que se encargará de respondera /.

func PaginaPrincipal(c *gin.Context) {
	c.String(200, "Bienvenido a la página principal")
}

//Handler que se encargará de verificar si el id que pasa el cliente existe en la base de datos

func BuscarEmpleado(c *gin.Context) {
	empleado, ok := empleados[c.Param("id")]

	if ok {
		c.String(200, "El empleado con id %s se llama %s", c.Param("id"), empleado)
	} else {
		c.String(200, "El empleado con id %s no existe", c.Param("id"))
	}

}
