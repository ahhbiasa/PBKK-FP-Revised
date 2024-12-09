package repository

import (
	entities "PBKK-FP-Revised/entities"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type CategoryRepository interface {
	Save(category entities.Category)
	Update(category entities.Category)
	Delete(category entities.Category)
	FindAll() []entities.Category
	FindByID(id int) (entities.Category, error)
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func New() CategoryRepository {
	// Define the DSN
	dsn := "root:@tcp(127.0.0.1:3306)/go-crud?charset=utf8mb4&parseTime=True&loc=Local"

	// Connect to the database using GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Println("Connected to the database using GORM")
	DB = db
	return &database{connection: db} // Add this return statement
}

func (db *database) CloseDB() {
	dbConn, err := db.connection.DB()
	if err != nil {
		panic(err)
	}
	dbConn.Close()
}

func (db *database) Save(category entities.Category) {
	db.connection.Create(&category)
}

func (db *database) Update(category entities.Category) {
	db.connection.Save(&category)
}

func (db *database) Delete(category entities.Category) {
	db.connection.Delete(&category)
}

func (db *database) FindAll() []entities.Category {
	var categories []entities.Category
	db.connection.Set("gorm:auto_preload", true).Find(&categories)
	return categories
}

// FindByID retrieves a category by its ID
func (db *database) FindByID(id int) (entities.Category, error) {
	var category entities.Category
	if err := db.connection.First(&category, id).Error; err != nil {
		return category, err // Return error if category is not found
	}
	return category, nil
}
