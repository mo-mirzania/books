package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mo-mirzania/books/services"
)

// GetBooks func
func GetBooks(c *gin.Context) {
	var bs services.BookServiceInterface = &services.BookService{}
	books, err := bs.GetBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, books)
}
