package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	productNames := [3]string{"Product 1", "Product 2", "Product 3"}
	prices := [3]float64{10.0, 20.0, 30.0}

	fmt.Println(productNames)
	fmt.Println(prices)

	for i, productName := range productNames {
		fmt.Printf("Product %d: %s - $%.2f\n", i+1, productName, prices[i])
	}

	featuredPrice := prices[1:3]
	featuredPrice[0] = 15.0
	fmt.Println("Featured Products:", featuredPrice)
	fmt.Println("Original Prices:", prices)

	productNamesDynamic := []string{"Product 1", "Product 2", "Product 3"}
	pricesDynamic := []float64{10.0, 20.0, 30.0}

	fmt.Printf("%p", &productNamesDynamic)
	fmt.Println(productNamesDynamic)
	fmt.Println(pricesDynamic)

	productNamesDynamic = append(productNamesDynamic, "Product 4")
	pricesDynamic = append(pricesDynamic, 40.0)

	fmt.Printf("%p", &productNamesDynamic)
	fmt.Println(productNamesDynamic)
	fmt.Println(pricesDynamic)
}
