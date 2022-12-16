package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	tx := db.Begin()
	var c Category
	err = tx.Debug().Clauses(clause.Loking{Strength: "UPDATE"}).First(&c, 1).Error
	if err != nil {
		panic(err)
	}
	c.Name = "Eletronic"
	tx.Debug().Save(&c)
	tx.Commit()
}
