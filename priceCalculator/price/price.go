package price

import (
	"example.com/price-calculator/helper"
	"example.com/price-calculator/ioManager"
)

type Price struct {
	Prices        []int64             `json:"prices"`
	TaxRate       int64               `json:"taxRate"`
	PricesWithTax []int64             `json:"pricesWithTax"`
	IOManager     ioManager.IOManager `json:"-"`
}

func New(ioManager ioManager.IOManager) (Price, error) {
	var data Price
	err := ioManager.Get(&data)

	if err != nil {
		return Price{}, err
	}

	return Price{
		Prices:        data.Prices,
		TaxRate:       data.TaxRate,
		PricesWithTax: data.PricesWithTax,
		IOManager:     ioManager,
	}, nil
}

func (p *Price) CalculatePricesWithTax(taxRate int64) {
	pricesWithTax := make([]int64, len(p.Prices))

	for i, price := range p.Prices {
		pricesWithTax[i] = helper.ApplyTax(price, taxRate)
	}

	p.TaxRate = taxRate
	p.PricesWithTax = pricesWithTax
}

func (p *Price) Save(fileName string) error {
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

	return p.IOManager.WriteResult(fileName, formatedData)
}
