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

	errCh := make(chan error, len(taxRates))

	p := price.New(prices)

	for _, taxRate := range taxRates {
		p.CalculatePricesWithTax(taxRate)
		formatedData := p.FormatData()

		fileName := fmt.Sprintf("prices_%d.json", taxRate)

		go ps.IO.WriteResult(fileName, formatedData, errCh)

	}

	for range taxRates {
		err := <-errCh
		if err != nil {
			return err
		}
	}

	return nil
}
