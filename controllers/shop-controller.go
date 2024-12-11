package controllers

import (
	entities "PBKK-FP-Revised/entities"
	"PBKK-FP-Revised/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ShopController interface {
	FindAll() ([]entities.Shop, error) // Update to include error
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
	EditShop(ctx *gin.Context) error
}

type shopController struct {
	service service.ShopService
}

func NewShopController(service service.ShopService) ShopController {
	return &shopController{
		service: service,
	}
}

func (c *shopController) FindAll() ([]entities.Shop, error) {
	shop, err := c.service.FindAll()
	if err != nil {
		return nil, err
	}
	return shop, nil
}

func (c *shopController) Save(ctx *gin.Context) error {
	var shop entities.Shop
	err := ctx.ShouldBind(&shop)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	// Set created_at to the current time
	shop.CreatedAt = time.Now()
	shop.UpdatedAt = time.Now()

	err = c.service.Save(shop)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	return nil
}

func (c *shopController) ShowAll(ctx *gin.Context) {
	shops, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.HTML(http.StatusOK, "indexshop.html", gin.H{"shops": shops})
}

func (c *shopController) Update(ctx *gin.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return err
	}

	var updatedShop entities.Shop
	if err := ctx.ShouldBind(&updatedShop); err != nil {
		return err
	}

	// Set the ID
	updatedShop.ID = id

	// Do not modify CreatedAt during update
	var originalShop entities.Shop
	originalShop, err = c.service.GetShopByID(id)
	if err != nil {
		return err
	}

	// Keep the original CreatedAt value intact during the update
	updatedShop.CreatedAt = originalShop.CreatedAt

	// Proceed with the update
	c.service.Update(updatedShop)
	return nil
}

func (c *shopController) Delete(ctx *gin.Context) error {
	var shop entities.Shop
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	shop.ID = int(id)

	c.service.Delete(shop)
	return nil
}

func (c *shopController) EditShop(ctx *gin.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid shop ID"})
		return err
	}

	shop, err := c.service.GetShopByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Shop not found"})
		return err
	}

	// Pass the shop as a key-value pair to the template
	data := gin.H{
		"shop": shop, // Match this key with the template
	}

	ctx.HTML(http.StatusOK, "editshop.html", data)
	return nil
}

// func EditShopHandler(w http.ResponseWriter, r *http.Request) {
//     id := r.URL.Query().Get("id")
//     shop, err := shopService.GetShopByID(id)
//     if err != nil {
//         http.Error(w, "Shop not found", http.StatusNotFound)
//         return
//     }

//     tmpl, err := template.ParseFiles("templates/editshop.html")
//     if err != nil {
//         http.Error(w, "Internal Server Error", http.StatusInternalServerError)
//         return
//     }

//     data := struct {
//         Shop Shop
//     }{
//         Shop: shop,
//     }

//     tmpl.Execute(w, data)
// }
