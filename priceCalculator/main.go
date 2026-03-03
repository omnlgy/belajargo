package main

import (
	"fmt"
	"log"

	"example.com/price-calculator/price"
	"example.com/price-calculator/tax"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const taxScale int64 = 10000
const moneyScale int64 = 100

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	priceJson, taxJson, err := getPriceAndTax()

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("Hello, World!")

	prices := priceJson.Prices
	taxRates := taxJson.Rates

	result := make(map[string][]string, len(taxRates))

	for _, taxRate := range taxRates {
		listPriceWithTax := make([]string, len(prices))

		for priceIndex, price := range prices {
			priceWithTax := applyTax(price, taxRate)

			listPriceWithTax[priceIndex] = formatRupiah(priceWithTax)
		}

		stringPrecentage := bpsFormat(taxRate)

		result[stringPrecentage] = listPriceWithTax
	}

	fmt.Println(result)
}

func getPriceAndTax() (*price.Price, *tax.TaxRate, error) {
	priceJson, err := price.Get()

	if err != nil {
		return nil, nil, fmt.Errorf("Error get prices %w", err)
	}

	taxJson, err := tax.Get()

	if err != nil {
		return nil, nil, fmt.Errorf("Error get tax %w", err)
	}

	return &priceJson, &taxJson, nil

}

func applyTax(price int64, taxRate int64) int64 {
	return price * (taxScale + taxRate) / taxScale
}

func formatRupiah(amount int64) string {
	p := message.NewPrinter(language.Indonesian)

	rupiah := amount / moneyScale
	sen := amount % moneyScale

	return p.Sprintf("Rp %d,%02d", rupiah, sen)
}

func bpsFormat(taxRate int64) string {
	integer := taxRate / 100
	decimal := taxRate % 100

	return fmt.Sprintf("%d.%.2d%%", integer, decimal)
}
