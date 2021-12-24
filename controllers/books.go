package controllers

import (
	"net/http"

	"toy/models"

	"github.com/gin-gonic/gin"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// FindBooks godoc
// @Summary Find all books
// @Description Find all books
// @Tags         book
// @Accept       json
// @Produce      json
// @Success 200 {object} []models.Book
// @Router /books [get]
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// FindBook godoc
// @Summary Find book
// @Description Find a book by ID
// @Tags         book
// @Id	 		 get-string-by-int
// @Accept       json
// @Produce      json
// @Param id path string true "Book Id"
// @Success 200 {object} models.Book
// @Router /book/:id [get]
func FindBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// CreateBook godoc
// @Summary Create a book
// @Description Create new book
// @Tags         book
// @Accept       json
// @Produce      json
// @Success 200 {object} models.Book
// @Router /book [post]
func CreateBook(c *gin.Context) {
	// Validate input
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// UpdateBook godoc
// @Summary 	 Update a book
// @Description  Update a book
// @Tags         book
// @Accept       json
// @Produce      json
// @Success 200 {object} models.Book
// @Router /book [patch]
func UpdateBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DeleteBook godoc
// @Summary delete a book
// @Description delete a book
// @Tags         book
// @Accept       json
// @Produce      json
// @Success 200 {object} boolean
// @Router /book [delete]
func DeleteBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
