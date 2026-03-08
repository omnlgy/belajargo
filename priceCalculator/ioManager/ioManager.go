package ioManager

type IOManager interface {
	// Get reads the input file or command line arguments and returns the prices.
	Get() ([]int64, error)
	// WriteResult writes the formatted data to a file or displayed on the comand line.
	WriteResult(fileName string, v any, chErr chan error)
}
