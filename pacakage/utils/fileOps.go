package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloadToFile(value float64, filename string) {
	valueString := fmt.Sprint(value)

	os.WriteFile("balance.txt", []byte(valueString), 0644)
}

func ReadFloadFromFile(fileName string) (float64, error) {
	valueByte, err := os.ReadFile(fileName)

	if err != nil {
		return 0, errors.New("balance file not found")
	}

	valueString := string(valueByte)
	valueFloat, err := strconv.ParseFloat(valueString, 64)

	if err != nil {
		return 0, errors.New("invalid balance format")
	}

	return valueFloat, nil
}
