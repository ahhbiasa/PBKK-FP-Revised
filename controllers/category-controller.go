package controllers

import (
	entities "PBKK-FP-Revised/entities"
	"PBKK-FP-Revised/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	FindAll() []entities.Category
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

func (c *controller) FindAll() []entities.Category {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var category entities.Category
	err := ctx.ShouldBind(&category)
	if err != nil {
		return err
	}
	c.service.Save(category)
	return nil
}

func (c *controller) ShowAll(ctx *gin.Context) {
	// Fetch all categories from the service
	categories := c.service.FindAll()

	// Prepare the data to pass to the template
	data := gin.H{
		"name":       "Category Name", // This could be a static value or a dynamic one if needed
		"categories": categories,      // List of categories to render in the template
	}

	// Render the HTML template and pass the data
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

	updatedCategory.ID = id
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
