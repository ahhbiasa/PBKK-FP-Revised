package controllers

import (
	entities "PBKK-FP-Revised/entities"
	"PBKK-FP-Revised/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ProductController interface {
	FindAll() ([]entities.Product, error) // Update to include error
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
	EditProduct(ctx *gin.Context) error
	AddProduct(ctx *gin.Context) error
}

type productController struct {
	service service.ProductService
}

func NewProductController(service service.ProductService) ProductController {
	return &productController{
		service: service,
	}
}

func (c *productController) FindAll() ([]entities.Product, error) {
	product, err := c.service.FindAll()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (c *productController) Save(ctx *gin.Context) error {
	var product entities.Product
	err := ctx.ShouldBind(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	// Set created_at to the current time
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	err = c.service.Save(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	return nil
}

func (c *productController) ShowAll(ctx *gin.Context) {
	products, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Ensure `products` is not nil
	if products == nil {
		products = []entities.Product{}
	}

	ctx.HTML(http.StatusOK, "indexproduct.html", gin.H{"products": products})
}

func (c *productController) Update(ctx *gin.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return err
	}

	var updatedProduct entities.Product
	if err := ctx.ShouldBind(&updatedProduct); err != nil {
		return err
	}

	// Set the ID
	updatedProduct.ID = id

	// Do not modify CreatedAt during update
	var originalProduct entities.Product
	originalProduct, err = c.service.GetProductByID(id)
	if err != nil {
		return err
	}

	// Keep the original CreatedAt value intact during the update
	updatedProduct.CreatedAt = originalProduct.CreatedAt

	// Proceed with the update
	c.service.Update(updatedProduct)
	return nil
}

func (c *productController) Delete(ctx *gin.Context) error {
	var product entities.Product
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	product.ID = int(id)

	c.service.Delete(product)
	return nil
}

func (c *productController) EditProduct(ctx *gin.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return err
	}

	// Fetch the product by ID
	product, err := c.service.GetProductByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return err
	}

	// Fetch the category associated with the product
	categories, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return err
	}

	// Fetch the shop associated with the product
	shops, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Shop not found"})
		return err
	}

	// Pass all details to the template
	data := gin.H{
		"product":    product,
		"categories": categories,
		"shops":      shops,
	}

	ctx.HTML(http.StatusOK, "editproduct.html", data)
	return nil
}

func (c *productController) AddProduct(ctx *gin.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return err
	}

	// Fetch the product by ID
	product, err := c.service.GetProductByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return err
	}

	// Fetch the category associated with the product
	category, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return err
	}

	// Fetch the shop associated with the product
	shop, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Shop not found"})
		return err
	}

	// Pass all details to the template
	data := gin.H{
		"product":  product,
		"category": category,
		"shop":     shop,
	}

	ctx.HTML(http.StatusOK, "editproduct.html", data)
	return nil
}
