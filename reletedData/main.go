package main

import (
	"fmt"
	"time"
)

type product struct {
	id        int
	title     string
	createdAt string
}

type products struct {
	products []product
}

func (p *products) addProducts(prductName string) {
	p.products = append(p.products, product{
		id:        len(p.products) + 1,
		title:     prductName,
		createdAt: time.Now().String(),
	})
}

func newProducts() products {
	return products{
		products: []product{},
	}
}

func main() {
	hobbies := [3]string{"bicycle", "fotography", "hiking"}

	fmt.Println("All of hobbies: ", hobbies)
	fmt.Println("One of my hobby: ", hobbies[0])
	fmt.Println("Rest of my hobbies: ", hobbies[1:])

	hobbies[0] = "running"
	fmt.Println("All of hobbies: ", hobbies)

	goals := []string{"swtich career", "learn go"}

	fmt.Println("All of goals: ", goals)

	goals = append(goals, "add new goal")
	fmt.Println("All of goals: ", goals)
	println("============================")
	products := newProducts()
	products.addProducts("Product 1")
	products.addProducts("Product 2")
	products.addProducts("Product 3")

	products.products[1].title = "Updated Product 2"

	fmt.Println("All of products: ", products)

}

// func main() {
// 	fmt.Println("Hello world")
// 	productNames := [3]string{"Product 1", "Product 2", "Product 3"}
// 	prices := [3]float64{10.0, 20.0, 30.0}

// 	fmt.Println(productNames)
// 	fmt.Println(prices)

// 	for i, productName := range productNames {
// 		fmt.Printf("Product %d: %s - $%.2f\n", i+1, productName, prices[i])
// 	}

// 	featuredPrice := prices[1:3]
// 	featuredPrice[0] = 15.0
// 	fmt.Println("Featured Products:", featuredPrice)
// 	fmt.Println("Original Prices:", prices)

// 	productNamesDynamic := []string{"Product 1", "Product 2", "Product 3"}
// 	pricesDynamic := []float64{10.0, 20.0, 30.0}

// 	fmt.Printf("%p", &productNamesDynamic)
// 	fmt.Println(productNamesDynamic)
// 	fmt.Println(pricesDynamic)

// 	productNamesDynamic = append(productNamesDynamic, "Product 4")
// 	pricesDynamic = append(pricesDynamic, 40.0)

// 	fmt.Printf("%p", &productNamesDynamic)
// 	fmt.Println(productNamesDynamic)
// 	fmt.Println(pricesDynamic)
// }
