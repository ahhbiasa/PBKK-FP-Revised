package repository

import (
	entities "PBKK-FP-Revised/entities"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ProductDB *gorm.DB

type ProductRepository interface {
	Save(product entities.Product) error
	Update(product entities.Product)
	Delete(product entities.Product)
	FindAll() ([]entities.Product, error)
	FindByID(id int) (entities.Product, error)
	CloseDB()
}

type productDatabase struct {
	connection *gorm.DB
}

func NewProductRepository() ProductRepository {
	dsn := "root:Xadenth04*@tcp(localhost:3306)/go_crud?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err) // Provide more details
	}
	log.Println("Connected to the database successfully")
	DB = db
	return &productDatabase{connection: db}
}

func (db *productDatabase) CloseDB() {
	dbConn, err := db.connection.DB()
	if err != nil {
		panic(err)
	}
	dbConn.Close()
}

func (db *productDatabase) Save(product entities.Product) error {
	result := db.connection.Create(&product)
	return result.Error
}

func (db *productDatabase) Update(product entities.Product) {
	db.connection.Save(&product)
}

func (db *productDatabase) Delete(product entities.Product) {
	db.connection.Delete(&product)
}

func (db *productDatabase) FindAll() ([]entities.Product, error) {
	var products []entities.Product
	result := db.connection.Preload("Category").Preload("Shop").Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	// Return an empty slice instead of nil
	return products, nil
}

// FindByID retrieves a category by its ID
func (db *productDatabase) FindByID(id int) (entities.Product, error) {
	var product entities.Product
	if err := db.connection.First(&product, id).Error; err != nil {
		return product, err
	}
	return product, nil
}
