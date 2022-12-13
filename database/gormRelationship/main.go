package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Category{})
	db.AutoMigrate(&Product{})
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
