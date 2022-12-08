package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func InsertProduct(db *sql.DB, product *Product) error {
	// for db sqlite ($1, $2, $3)
	stmt, err := db.Prepare("insert into products (id, name, price) values (?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func UpdateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ? , price = ? where id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func FindProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	var p Product
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func FindAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select id, name, price from products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func DeleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("Delete from products where id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	product := NewProduct("Monitor", 2000.00)
	err = InsertProduct(db, product)
	if err != nil {
		panic(err)
	}

	fmt.Print(product)
	product.Price = 250.00
	product.Name = "Teclado"

	err = UpdateProduct(db, product)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Updated product: %v", product)

	p, err := FindProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Find Product: %v\n", p)
	products, err := FindAllProducts(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("Result of all products:")
	for _, product := range products {
		fmt.Printf("Product Name: %s, Price: %.2f\n", product.Name, product.Price)
	}

	err = DeleteProduct(db, p.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deleted Product: %v\n", p.Name)
}
