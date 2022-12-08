package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ProductsGorm struct {
	ID int `gorm:"primaryKey"`
	Name string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}	
	db.AutoMigrate(&ProductsGorm{})
	// Crate a one Product
	// db.Create(&ProductsGorm{
	// 	Name: "Violao",
	// 	Price: 5000.00,
	// })
	products := []ProductsGorm{
		{Name: "Notebook", Price: 15000.00},
		{Name: "Monitor", Price: 2000.00},
		{Name: "Mouse", Price: 200.00},
		{Name: "Teclado", Price: 400.00},
	}
	db.Create(products)
}
