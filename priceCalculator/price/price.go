package price

import (
	"encoding/json"
	"os"
)

type Price struct {
	Prices []int64 `json:"prices"`
}

func Get() (Price, error) {
	data, err := os.ReadFile("./saveFiles/prices.json")

	if err != nil {
		return Price{}, err
	}

	var result Price

	err = json.Unmarshal(data, &result)

	if err != nil {
		return Price{}, err
	}

	return Price{
		Prices: result.Prices,
	}, nil
}
