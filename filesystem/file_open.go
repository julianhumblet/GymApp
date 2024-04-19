package filesystem

import "os"

func OpenFile(pathFile string) (*os.File, error) {
	file, err := os.Open(pathFile)
	if err != nil {
		return nil, err
	}
	return file, nil
}
