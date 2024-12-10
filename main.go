package main

import (
	"PBKK-FP-Revised/controllers"
	"PBKK-FP-Revised/repository"
	"PBKK-FP-Revised/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	CategoryRepository repository.CategoryRepository  = repository.New()
	CategoryService    service.CategoryService        = service.New(CategoryRepository)
	CategoryController controllers.CategoryController = controllers.New(CategoryService)
)

func main() {
	defer CategoryRepository.CloseDB()

	router := gin.Default()

	router.LoadHTMLGlob("templates/*.html")

	router.Static("/static", "./static")

	apiRoutes := router.Group("/api")
	{
		apiRoutes.GET("/categories", func(ctx *gin.Context) {
			categories, err := CategoryController.FindAll() // Handle the error properly
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(200, categories) // Return the categories as JSON
		})

		apiRoutes.POST("/categories", func(ctx *gin.Context) {
			err := CategoryController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Category created successfully"})
			}
		})

		apiRoutes.PUT("/categories/:id", func(ctx *gin.Context) {
			err := CategoryController.Update(ctx)
			if err != nil {
				// ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
			}
		})

		apiRoutes.DELETE("/categories/:id", func(ctx *gin.Context) {
			err := CategoryController.Delete(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
			}
		})
	}

	viewRoutes := router.Group("/view")
	{
		viewRoutes.GET("/categories", CategoryController.ShowAll)

		// Render the "Add Category" page
		viewRoutes.GET("/categories/add", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "createcategory.html", nil)
		})

		// Handle form submission for creating a category
		viewRoutes.POST("/categories/add", func(ctx *gin.Context) {
			err := CategoryController.Save(ctx)
			if err != nil {
				ctx.HTML(http.StatusBadRequest, "createcategory.html", gin.H{"error": err.Error()})
			} else {
				ctx.Redirect(http.StatusFound, "/view/categories")
			}
		})

		viewRoutes.GET("/categories/edit/:id", func(ctx *gin.Context) {
			CategoryController.EditCategory(ctx)
		})

		viewRoutes.GET("/", homePageHandler)
	}

	router.Run(":8080")
}

// Handler for rendering the home page
func homePageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}
