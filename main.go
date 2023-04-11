package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*Ejercicio 1 : Iniciando el proyecto
Debemos crear un repositorio en github.com para poder subir nuestros avances. Este repositorio es el que vamos a utilizar para
llevar lo que realicemos durante las distintas prácticas de Go Web.
Primero debemos clonar el repositorio creado, luego iniciar nuestro proyecto de go con con el comando go mod init.
El siguiente paso será crear un archivo main.go donde deberán cargar en una slice, desde un archivo JSON, los datos de productos.
Esta slice se debe cargar cada vez que se inicie la API para realizar las distintas consultas.
*/

type Products struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Quantity       int     `json:"quantity"`
	CodeValue      string  `json:"code_value"`
	IsPublished    bool    `json:"is_published"`
	ExpirationDate string  `json:"expiration_date"`
	Price          float64 `json:"price"`
}

// func generateMap() map[int]Products {
// 	products, err := LoadProducts("products.json")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	productsMap := make(map[int]Products)
// 	for _, product := range products {
// 		productsMap[product.ID] = product
// 	}
// 	return productsMap
// }

// Cargamos los productos desde el archivo JSON en un slice de productos structc
func LoadProducts(filename string) ([]Products, error) {
	var products []Products
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&products)

	if err != nil {
		return nil, err
	}

	return products, nil
}

// Busqueda de producto por ID. Utilizamos LoadProducts para cargar los productos desde el archivo JSON.
func FindProduct(c *gin.Context) {
	products, err := LoadProducts("products.json")
	if err != nil {
		log.Fatal(err)
	}

	//Pasamos parametro convertiendo a int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Recorremos slice de productos y comparamos el id pasado por parametro con el id de cada producto. Diferente logia al map de clave valor del ejemplo
	for _, p := range products {
		if p.ID == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})

}

func FilterProductsByPrice(productos []Products, priceF float64) []Products {

	var result []Products
	for _, p := range productos {
		if p.Price > priceF {
			result = append(result, p)
		}
	}
	return result
}

func main() {
	products, err := LoadProducts("products.json")
	if err != nil {
		log.Fatal(err)
	}

	server := gin.Default()

	//Crear una ruta /ping que debe respondernos con un string que contenga pong con el status 200 OK.
	server.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	//Crear una ruta /products que nos devuelva la lista de todos los productos en la slice.
	server.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, products)
	})

	//Crear una ruta /products/:id que nos devuelva el producto que tenga el id que se pasa por parámetro.
	server.GET("/products/:id", FindProduct)

	//Crear una ruta /products/search que nos permita buscar por parámetro los productos cuyo precio sean mayor a un valor priceGt.
	server.GET("/products/search", func(c *gin.Context) {

		priceFStr := c.Query("priceF")

		priceF, err := strconv.ParseFloat(priceFStr, 64)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		filtredProducts := FilterProductsByPrice(products, 995)
		c.JSON(http.StatusOK, filtredProducts)
	})

	server.Run(":8080")
}
