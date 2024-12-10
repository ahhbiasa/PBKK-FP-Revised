package controllers

import (
	entities "PBKK-FP-Revised/entities"
	"PBKK-FP-Revised/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	FindAll() ([]entities.Category, error) // Update to include error
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
	EditCategory(ctx *gin.Context) error
}

type controller struct {
	service service.CategoryService
}

// New creates a new CategoryController with the given service.
func New(service service.CategoryService) CategoryController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() ([]entities.Category, error) {
	categories, err := c.service.FindAll() // Get categories from the service
	if err != nil {
		return nil, err // Return nil categories and the error
	}
	return categories, nil // Return the categories and nil error
}

func (c *controller) Save(ctx *gin.Context) error {
	var category entities.Category
	err := ctx.ShouldBind(&category)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	// Set created_at to the current time
	category.CreatedAt = time.Now()

	savedCategory, err := c.service.Save(category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":  "Category created successfully",
		"category": savedCategory,
	})
	return nil
}

func (c *controller) ShowAll(ctx *gin.Context) {
	categories, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	data := gin.H{
		"name":       "Category Name",
		"categories": categories,
	}

	ctx.HTML(http.StatusOK, "indexcategories.html", data)
}

func (c *controller) Update(ctx *gin.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return err
	}

	var updatedCategory entities.Category
	if err := ctx.ShouldBind(&updatedCategory); err != nil {
		return err
	}

	// Set the ID
	updatedCategory.ID = id

	// Do not modify CreatedAt during update
	var originalCategory entities.Category
	originalCategory, err = c.service.GetCategoryByID(id) // Fetch the original category
	if err != nil {
		return err
	}

	// Keep the original CreatedAt value intact during the update
	updatedCategory.CreatedAt = originalCategory.CreatedAt

	// Proceed with the update
	c.service.Update(updatedCategory)
	ctx.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
	return nil
}

func (c *controller) Delete(ctx *gin.Context) error {
	var category entities.Category
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	category.ID = int(id)

	c.service.Delete(category)
	return nil
}

// EditCategory fetches a category by its ID and renders it on the edit page
func (c *controller) EditCategory(ctx *gin.Context) error {
	id, err := strconv.Atoi(ctx.Param("id")) // Parse the category ID from the URL parameter
	if err != nil {
		return err
	}

	// Fetch the category by ID from the service
	category, err := c.service.GetCategoryByID(id)
	if err != nil {
		return err
	}

	// Render the category data into the template
	data := gin.H{
		"category": category,
	}

	// Render the template with the category data for editing
	ctx.HTML(http.StatusOK, "editcategory.html", data)
	return nil
}
