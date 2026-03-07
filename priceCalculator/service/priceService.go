package service

import (
	"fmt"

	"example.com/price-calculator/ioManager"
	"example.com/price-calculator/price"
)

type PriceService struct {
	IO ioManager.IOManager
}

func NewPriceService(io ioManager.IOManager) PriceService {
	return PriceService{
		IO: io,
	}
}

func (ps *PriceService) Run(taxRates []int64) error {
	prices, err := ps.IO.Get()

	if err != nil {
		return err
	}

	p := price.New(prices)

	for _, taxRate := range taxRates {

		p.CalculatePricesWithTax(taxRate)

		formatedData := p.FormatData()

		err := ps.IO.WriteResult(fmt.Sprintf("prices_%d.json", taxRate), formatedData)

		if err != nil {
			return fmt.Errorf("Error writing file %w", err)
		}
	}

	return nil
}
