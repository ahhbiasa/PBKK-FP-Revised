package main

import (
	"PBKK-FP-Revised/controllers"
	"PBKK-FP-Revised/repository"
	"PBKK-FP-Revised/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	CategoryRepository repository.CategoryRepository  = repository.NewCategoryRepository()
	CategoryService    service.CategoryService        = service.NewCategoryController(CategoryRepository)
	CategoryController controllers.CategoryController = controllers.NewCategoryController(CategoryService)

	ProductRepository repository.ProductRepository  = repository.NewProductRepository()
	ProductService    service.ProductService        = service.NewProductService(ProductRepository, CategoryRepository, ShopRepository)
	ProductController controllers.ProductController = controllers.NewProductController(ProductService)

	ShopRepository repository.ShopRepository  = repository.NewShopRepository()
	ShopService    service.ShopService        = service.NewShopService(ShopRepository)
	ShopController controllers.ShopController = controllers.NewShopController(ShopService)
)

func main() {
	defer CategoryRepository.CloseDB()
	defer ShopRepository.CloseDB()

	router := gin.Default()

	router.LoadHTMLGlob("templates/*.html")

	router.Static("/static", "./static")

	apiRoutes := router.Group("/api")
	{

		apiRoutes.POST("/categories/add", func(ctx *gin.Context) {
			err := CategoryController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.Redirect(http.StatusFound, "/view/categories")
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

		apiRoutes.POST("/products/add", func(ctx *gin.Context) {
			err := ProductController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Category created successfully"})
			}
		})

		apiRoutes.PUT("/products/:id", func(ctx *gin.Context) {
			err := ProductController.Update(ctx)
			if err != nil {
				// ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
			}
		})

		apiRoutes.DELETE("/products/:id", func(ctx *gin.Context) {
			err := ProductController.Delete(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
			}
		})

		apiRoutes.POST("/shops/add", func(ctx *gin.Context) {
			err := ShopController.Save(ctx)
			if err != nil {
				ctx.HTML(http.StatusBadRequest, "createshop.html", gin.H{"error": err.Error()})
			} else {
				ctx.Redirect(http.StatusFound, "/view/shops")
			}
		})

		apiRoutes.PUT("/shops/:id", func(ctx *gin.Context) {
			err := ShopController.Update(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Shop updated successfully"})
			}
		})

		apiRoutes.DELETE("/shops/:id", func(ctx *gin.Context) {
			err := ShopController.Delete(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Shop deleted successfully"})
			}
		})
	}

	viewRoutes := router.Group("/view")
	{
		viewRoutes.GET("/categories", CategoryController.ShowAll)

		viewRoutes.GET("/categories/add", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "createcategory.html", nil)
		})

		viewRoutes.GET("/categories/edit/:id", func(ctx *gin.Context) {
			CategoryController.EditCategory(ctx)
		})

		viewRoutes.GET("/shops", ShopController.ShowAll)

		viewRoutes.GET("/shops/add", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "createshop.html", nil)
		})

		viewRoutes.GET("/shops/edit/:id", func(ctx *gin.Context) {
			ShopController.EditShop(ctx)
		})

		viewRoutes.GET("/products", ProductController.ShowAll)

		viewRoutes.GET("/products/add", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "createproduct.html", nil)
		})

		viewRoutes.GET("/products/edit/:id", func(ctx *gin.Context) {
			ProductController.EditProduct(ctx)
		})

		viewRoutes.GET("/", homePageHandler)
	}
	router.Run(":8080")
}

// Handler for rendering the home page
func homePageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}
