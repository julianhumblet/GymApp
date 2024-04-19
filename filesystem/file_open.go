package filesystem

import (
	"fmt"
	"os"
)

func OpenFile(pathFile string) (*os.File, error) {
	file, err := os.Open(pathFile)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}

	return file, nil
}
