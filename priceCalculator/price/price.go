package price

import (
	"example.com/price-calculator/helper"
)

type Price struct {
	Prices        []int64
	TaxRate       int64
	PricesWithTax []int64
}

func New(prices []int64) Price {
	return Price{
		Prices: prices,
	}
}

func (p *Price) CalculatePricesWithTax(taxRate int64) {
	pricesWithTax := make([]int64, len(p.Prices))

	for i, price := range p.Prices {
		pricesWithTax[i] = helper.ApplyTax(price, taxRate)
	}

	p.TaxRate = taxRate
	p.PricesWithTax = pricesWithTax
}

func (p *Price) FormatData() any {
	var formatedData struct {
		Prices        []string `json:"prices"`
		TaxRate       string   `json:"taxRate"`
		PricesWithTax []string `json:"pricesWithTax"`
	}

	formatedPrices := make([]string, len(p.Prices))
	formatedPricesWithTax := make([]string, len(p.PricesWithTax))

	for i, price := range p.Prices {
		formatedPrices[i] = helper.FormatRupiah(price)
	}

	for i, price := range p.PricesWithTax {
		formatedPricesWithTax[i] = helper.FormatRupiah(price)
	}

	formatedData.Prices = formatedPrices
	formatedData.TaxRate = helper.BpsFormat(p.TaxRate)
	formatedData.PricesWithTax = formatedPricesWithTax

	return formatedData
}
