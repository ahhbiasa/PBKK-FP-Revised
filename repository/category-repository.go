package repository

import (
	entities "PBKK-FP-Revised/entities"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type CategoryRepository interface {
	Save(category entities.Category) error
	Update(category entities.Category)
	Delete(category entities.Category)
	FindAll() ([]entities.Category, error)
	FindByID(id int) (entities.Category, error)
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func New() CategoryRepository {
	dsn := "root:Xadenth04*@tcp(localhost:3306)/go_crud?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err) // Provide more details
	}
	log.Println("Connected to the database successfully")
	DB = db
	return &database{connection: db}
}

func (db *database) CloseDB() {
	dbConn, err := db.connection.DB()
	if err != nil {
		panic(err)
	}
	dbConn.Close()
}

func (db *database) Save(category entities.Category) error {
	result := db.connection.Create(&category)
	return result.Error
}

func (db *database) Update(category entities.Category) {
	db.connection.Save(&category)
}

func (db *database) Delete(category entities.Category) {
	db.connection.Delete(&category)
}

func (db *database) FindAll() ([]entities.Category, error) {
	var categories []entities.Category
	result := db.connection.Set("gorm:auto_preload", true).Find(&categories)
	return categories, result.Error
}

// FindByID retrieves a category by its ID
func (db *database) FindByID(id int) (entities.Category, error) {
	var category entities.Category
	if err := db.connection.First(&category, id).Error; err != nil {
		return category, err // Return error if category is not found
	}
	return category, nil
}
