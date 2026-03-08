package cmdmanager

import (
	"encoding/json"
	"fmt"
)

type CmdManager struct{}

func New() CmdManager {
	return CmdManager{}
}

func (c CmdManager) Get() ([]int64, error) {
	var price int64
	var prices []int64

	for {
		fmt.Print("Input Price: ")
		fmt.Scan(&price)
		if price <= 0 {
			fmt.Println("Price must be positive")
			break
		}

		prices = append(prices, price)
	}

	return prices, nil
}

func (c CmdManager) WriteResult(_ string, v any, chErr chan error) {

	jsonData, err := json.Marshal(v)

	if err != nil {
		chErr <- err
		return
	}

	fmt.Println("result:", string(jsonData))

}
