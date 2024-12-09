package controllers

import (
	entities "PBKK-FP-Revised/entities"
	"PBKK-FP-Revised/service"
	"fmt"
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
}

type controller struct {
	service service.CategoryService
}

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
	category := c.service.FindAll()
	fmt.Println("Categories fetched:", category) // Debugging line to see what is fetched
	data := gin.H{
		"name":       "Category Name",
		"categories": category,
	}
	ctx.HTML(http.StatusOK, "indexcategories.html", data)
}

func (c *controller) Update(ctx *gin.Context) error {
	var category entities.Category
	err := ctx.ShouldBindJSON(&category)
	if err != nil {
		return err
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	category.ID = int(id)

	c.service.Update(category)
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
