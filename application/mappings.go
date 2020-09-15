package application

import "github.com/mo-mirzania/books/controllers"

// Mappings func
func Mappings() {
	router.GET("/books", controllers.GetBooks)
}
