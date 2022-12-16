package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
	// hasMany
	Products []Product
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

type ProductMany struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []CategoryMany `gorm:"many2many:products_categories"`
	gorm.Model
}

type CategoryMany struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []ProductMany `gorm:"many2many:products_categories"`
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// belongsTo(db)
	// hasOne(db)
	// hasMany(db)
	manyToMany(db)
}

func belongsTo(db *gorm.DB) {
	db.AutoMigrate(&Category{}, &Product{})
	db.Create(&Category{
		Name: "Electronic",
	})
	db.Create(&Product{
		Name:       "Notebook",
		Price:      16000.00,
		CategoryID: 1,
	})
	var products []Product
	db.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.ID, product.Category.Name)
	}
}

func hasOne(db *gorm.DB) {
	db.AutoMigrate(&Category{}, &Product{}, &SerialNumber{})
	db.Create(&Category{
		Name: "Electronic",
	})
	db.Create(&Product{
		Name:       "Notebook",
		Price:      16000.00,
		CategoryID: 1,
	})
	db.Create(&SerialNumber{
		Number:    "35264",
		ProductID: 1,
	})
	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	}
}

func hasMany(db *gorm.DB) {
	db.AutoMigrate(&Category{}, &Product{}, &SerialNumber{})
	db.Create(&Category{
		Name: "Electronic",
	})
	db.Create(&Product{
		Name:       "Notebook",
		Price:      16000.00,
		CategoryID: 1,
	})
	db.Create(&SerialNumber{
		Number:    "35264",
		ProductID: 1,
	})
	var categories []Category
	err := db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println(" - ", product.Name, product.SerialNumber.Number)
		}
	}
}

func manyToMany(db *gorm.DB) {
	db.AutoMigrate(&ProductMany{}, &CategoryMany{})
	category := CategoryMany{Name: "Electronic"}
	db.Create(&category)
	category2 := CategoryMany{Name: "kitchen"}
	db.Create(&category)
	db.Create(&ProductMany{
		Name:       "Panela",
		Price:      200.00,
		Categories: []CategoryMany{category, category2},
	})

	var categories []CategoryMany
	err := db.Model(&CategoryMany{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println(" - ", product.Name)
		}
	}
}
