package filesystem

import (
	"log"
	"os"
)

func FileExists(pathFile string) bool {
	_, err := os.Stat(pathFile)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}

	// Different error
	log.Printf("error checking file existence: %s", err)

	return false
}
