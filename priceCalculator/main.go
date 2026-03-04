package main

import (
	"fmt"
	"log"

	filemanager "example.com/price-calculator/fileManager"
	"example.com/price-calculator/price"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	taxRates := []int64{0, 100, 220, 3000, 45, 3400}
	fm := filemanager.New("./saveFiles/prices.json", "./saveFiles/")

	for _, taxRate := range taxRates {
		price, err := price.New(fm)

		if err != nil {
			log.Println(err)
			return
		}

		price.CalculatePricesWithTax(taxRate)

		err = price.Save(fmt.Sprintf("prices_%d.json", taxRate))

		if err != nil {
			log.Println(err)
			return
		}
	}
}
