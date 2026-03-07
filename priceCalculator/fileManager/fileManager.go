package filemanager

import (
	"encoding/json"
	"fmt"
	"os"
)

type FileManager struct {
	inputPath  string
	outputPath string
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		inputPath:  inputPath,
		outputPath: outputPath,
	}
}

func (f FileManager) Get() ([]int64, error) {
	data, err := os.ReadFile(f.inputPath)

	if err != nil {
		return nil, fmt.Errorf("Error read file %w", err)
	}

	var fileData struct {
		Prices []int64 `json:"prices"`
	}
	err = json.Unmarshal(data, &fileData)

	if err != nil {
		return nil, fmt.Errorf("Error unmarshal file %w", err)
	}

	return fileData.Prices, nil
}

func (f FileManager) WriteResult(fileName string, v any) error {
	jsonData, err := json.Marshal(v)

	if err != nil {
		return fmt.Errorf("Error marshal file %w", err)
	}

	err = os.WriteFile(f.outputPath+fileName, jsonData, 0644)

	if err != nil {
		return fmt.Errorf("Error write file %w", err)
	}

	return nil
}
