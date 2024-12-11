package repository

import (
	entities "PBKK-FP-Revised/entities"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ShopDB *gorm.DB

type ShopRepository interface {
	Save(shop entities.Shop) error
	Update(shop entities.Shop)
	Delete(shop entities.Shop)
	FindAll() ([]entities.Shop, error)
	FindByID(id int) (entities.Shop, error)
	CloseDB()
}

type shopDatabase struct {
	connection *gorm.DB
}

func NewShopRepository() ShopRepository {
	dsn := "root:Xadenth04*@tcp(localhost:3306)/go_crud?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err) // Provide more details
	}
	log.Println("Connected to the database successfully")
	DB = db
	return &shopDatabase{connection: db}
}

func (db *shopDatabase) CloseDB() {
	dbConn, err := db.connection.DB()
	if err != nil {
		panic(err)
	}
	dbConn.Close()
}

func (db *shopDatabase) Save(shop entities.Shop) error {
	result := db.connection.Create(&shop)
	return result.Error
}

func (db *shopDatabase) Update(shop entities.Shop) {
	db.connection.Save(&shop)
}

func (db *shopDatabase) Delete(shop entities.Shop) {
	db.connection.Delete(&shop)
}

func (db *shopDatabase) FindAll() ([]entities.Shop, error) {
	var shops []entities.Shop
	result := db.connection.Set("gorm:auto_preload", true).Find(&shops)
	if result.Error != nil {
		return nil, result.Error
	}

	if shops == nil {
		shops = []entities.Shop{}
	}
	return shops, nil
}

func (db *shopDatabase) FindByID(id int) (entities.Shop, error) {
	var shop entities.Shop
	if err := db.connection.First(&shop, id).Error; err != nil {
		return shop, err
	}
	return shop, nil
}
