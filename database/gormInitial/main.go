package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ProductsGorm struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

type ProductModel struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	// dsn := "root:root@tcp(localhost:3306)/goexpert"
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&ProductModel{})
	// db.Create(&ProductModel{
	// 	Name:  "Alexia",
	// 	Price: 400.00,
	// })
	// var product ProductModel
	// db.First(&product, 2)
	// product.Name = "Echo Dot"
	// db.Save(&product)
	// fmt.Println(product)
	// var product ProductModel
	// db.First(&product, 3)
	// db.Delete(&product)

	//table without the model
	// db.AutoMigrate(&ProductsGorm{})
	// Crate a one Product
	// db.Create(&ProductsGorm{
	// 	Name: "Violao",
	// 	Price: 5000.00,
	// })

	// products := []ProductsGorm{
	// 	{Name: "Notebook", Price: 15000.00},
	// 	{Name: "Monitor", Price: 2000.00},
	// 	{Name: "Mouse", Price: 200.00},
	// 	{Name: "Teclado", Price: 400.00},
	// }
	// db.Create(products)

	//select one
	// var product ProductsGorm
	// db.First(&product, 2)
	// fmt.Println(product)
	// db.First(&product, "name = ?", "Mouse")
	// fmt.Println(product)

	//select all
	// var products []ProductsGorm
	// db.Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	//Pagination and Limits
	// var products []ProductsGorm
	// db.Limit(3).Offset(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	//where
	// var products []ProductsGorm
	// db.Where("price > ?", 1000).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	//Update
	// var product ProductsGorm
	// db.First(&product, 2)
	// product.Name = "NooteBook AlienWare"
	// db.Save(&product)
	// fmt.Println(product)

	//Delete
	// var product ProductsGorm
	// db.First(&product, 3)
	// db.Delete(&product)
}
