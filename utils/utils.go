package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ScanFile(path string) (*bufio.Scanner, error) {

	file, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("FileToBufioReader: %w", err)
	}

	scan := bufio.NewScanner(file)

	return scan, nil
}
