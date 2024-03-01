package main

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       float32
	Description string
	Category    string
}

func main() {
	newProduct := Product{
		ID:          3,
		Name:        "Berenjena",
		Price:       1200,
		Description: "Berenjena",
		Category:    "Verdura",
	}

	newProduct.GetAll()
	newProduct.Save()
	newProduct.GetAll()
	product2, err := newProduct.getById(2)

	if err != "" {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(product2)
	}

}

var Products = []Product{
	{ID: 1, Name: "Tomate", Price: 1000, Description: "Tomate perita", Category: "Fruta"},
	{ID: 2, Name: "Banana", Price: 2000, Description: "Banana de ecuador", Category: "Fruta"}}

func (product Product) Save() {
	Products = append(Products, product)
}

func (Product) GetAll() {
	fmt.Printf("%v \n", Products)
}

func (Product) getById(ID int) (product Product, err string) {
	for _, current_product := range Products {
		if current_product.ID == ID {
			product = current_product
			return
		}
	}
	err = "No product with ID"
	return
}
