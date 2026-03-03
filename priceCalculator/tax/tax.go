package tax

import (
	"encoding/json"
	"os"
)

type TaxRate struct {
	Rates []int64 `json:"rates"`
}

type TaxResult struct {
	Percentage string
	Prices     []string
}

func Get() (TaxRate, error) {
	data, err := os.ReadFile("./saveFiles/taxRates.json")

	if err != nil {
		return TaxRate{}, err
	}

	var result TaxRate

	err = json.Unmarshal(data, &result)

	if err != nil {
		return TaxRate{}, err
	}

	return TaxRate{
		Rates: result.Rates,
	}, nil
}
