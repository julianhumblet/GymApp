package logfile

import (
	"GymApp/filesystem"
	"fmt"
	"log"
)

func LoadLogfile(pathLogfile string) error {
	if !filesystem.FileExists(pathLogfile) {
		return fmt.Errorf("logfile not found at: %s", pathLogfile)
	}

	logfile, err := filesystem.OpenFile(pathLogfile)
	if err != nil {
		return err
	}
	defer logfile.Close()

	log.SetOutput(logfile)

	return nil
}
