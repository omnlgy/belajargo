package main

import (
	"log"

	// cmdmanager "example.com/price-calculator/cmdManager"

	cmdmanager "example.com/price-calculator/cmdManager"
	"example.com/price-calculator/service"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	taxRates := []int64{0, 100, 220, 3000, 45, 3400}

	cm := cmdmanager.New()
	// fm := filemanager.New("./saveFiles/prices.json", "./saveFiles/")
	ps := service.NewPriceService(cm)

	err := ps.Run(taxRates)

	if err != nil {
		log.Fatal(err)
	}
}
