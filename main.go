package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Quantity       int     `json:"quantity"`
	CodeValue      string  `json:"code_value"`
	IsPublished    bool    `json:"is_published"`
	ExpirationDate string  `json:"expiration_date"`
	Price          float64 `json:"price"`
}

var productsList = []Product{}

type ProductRecuest struct {
	Name           string  `json:"name"`
	Quantity       int     `json:"quantity"`
	CodeValue      string  `json:"code_value"`
	IsPublished    bool    `json:"is_published"`
	ExpirationDate string  `json:"expiration_date"`
	Price          float64 `json:"price"`
}

// Cargamos los productos desde el archivo JSON en un slice de productos structc
func LoadProducts(path string, list *[]Product) {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal([]byte(file), &list)
	if err != nil {
		log.Fatal(err)
	}
}

// Busqueda de producto por ID. Utilizamos LoadProducts para cargar los productos desde el archivo JSON.
func FindProduct() gin.HandlerFunc {
	return func(c *gin.Context) {

		idParam := c.Param("id")
		//Pasamos parametro convertiendo a int
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//Recorremos slice de productos y comparamos el id pasado por parametro con el id de cada producto. Diferente logia al map de clave valor del ejemplo
		for _, p := range productsList {
			if p.ID == id {
				c.JSON(http.StatusOK, p)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
	}
}

// Filtramos productos por precio
func FilterProductsByPrice(productos []Product, priceF float64) []Product {

	var result []Product
	for _, p := range productos {
		if p.Price > priceF {
			result = append(result, p)
		}
	}
	return result
}

// Busqueda de productos por precio
func FindByPrice() gin.HandlerFunc {
	return func(c *gin.Context) {

		priceFStr := c.Query("priceF")

		priceF, err := strconv.ParseFloat(priceFStr, 64)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		filtredProducts := FilterProductsByPrice(productsList, priceF)
		c.JSON(http.StatusOK, filtredProducts)
	}
}

func FindProductAfterAdd(c *gin.Context) {

}

func AddProduct() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := len(productsList) + 1

		var newProduct ProductRecuest

		err2 := c.ShouldBindJSON(&newProduct)
		if err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
			return
		}

		if newProduct.Name == "" || newProduct.CodeValue == "" || newProduct.ExpirationDate == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No puede haber campos vacios"})
			return
		}

		//Verificamos que el codigo no exista
		for _, p := range productsList {
			if p.CodeValue == newProduct.CodeValue {
				c.JSON(http.StatusBadRequest, gin.H{"error": "El codigo ya existe"})
				return
			}
		}

		//Verificamos que la fecha sea valida
		_, err1 := time.Parse("02/01/2006", newProduct.ExpirationDate)
		if err1 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "La fecha no es valida"})
			return
		}

		//Creamos el producto
		prod := Product{
			ID:             id,
			Name:           newProduct.Name,
			Quantity:       newProduct.Quantity,
			CodeValue:      newProduct.CodeValue,
			IsPublished:    newProduct.IsPublished,
			ExpirationDate: newProduct.ExpirationDate,
			Price:          newProduct.Price,
		}

		//Agregamos el producto al slice
		productsList = append(productsList, prod)

		c.JSON(http.StatusOK, productsList)
		//c.JSON(http.StatusOK, prod)
	}
}

func main() {

	LoadProducts("products.json", &productsList)

	server := gin.Default()

	//Crear una ruta /ping que debe respondernos con un string que contenga pong con el status 200 OK.
	server.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	//Crear una ruta /products que nos devuelva la lista de todos los productos en la slice.
	server.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, productsList)
	})

	//Crear una ruta /products/:id que nos devuelva el producto que tenga el id que se pasa por parámetro.
	server.GET("/products/:id", FindProduct())

	//Crear una ruta /products/search que nos permita buscar por parámetro los productos cuyo precio sean mayor a un valor priceGt.
	server.GET("/products/search", FindByPrice())

	//Crear una ruta /products que nos permita agregar un producto a la slice, mediante POST
	server.POST("/products", AddProduct())

	//Utilizamos metodo de fila 183 para acceder al producto agregado
	// server.GET("/products/new/:id", func(c *gin.Context) {

	// 	//Pasamos parametro convertiendo a int
	// 	id, err := strconv.Atoi(c.Param("id"))
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	//Recorremos slice de productos y comparamos el id pasado por parametro con el id de cada producto. Diferente logia al map de clave valor del ejemplo
	// 	for _, p := range products {
	// 		if p.ID == id {
	// 			c.JSON(http.StatusOK, p)
	// 			return
	// 		}
	// 	}
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
	// })

	// pr := server.Group("/products")
	// pr.POST("/", AddProduct())
	// pr.GET("/:id", FindProduct)

	server.Run(":8080")
}
