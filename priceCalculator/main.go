package main

import (
	"log"

	// cmdmanager "example.com/price-calculator/cmdManager"
	filemanager "example.com/price-calculator/fileManager"
	"example.com/price-calculator/service"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	taxRates := []int64{0, 100, 220, 3000, 45, 3400}

	// cm := cmdmanager.New()
	fm := filemanager.New("./saveFiles/prices.json", "./saveFiles/")
	ps := service.NewPriceService(fm)

	err := ps.Run(taxRates)

	if err != nil {
		log.Fatal(err)
	}
}
