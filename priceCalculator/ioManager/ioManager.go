package ioManager

type IOManager interface {
	// Get reads the input file or command line arguments and unmarshals its contents into dest.
	// dest must be a pointer to the destination.
	Get(dest any) error
	// WriteResult writes the formatted data to a file.
	WriteResult(fileName string, v any) error
}
