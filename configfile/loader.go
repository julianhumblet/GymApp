package configfile

import (
	"GymApp/filesystem"
	"encoding/json"
	"fmt"
)

func LoadConfig(pathConfigfile string) (*Config, error) {

	if !filesystem.FileExists(pathConfigfile) {
		return nil, fmt.Errorf("config file not found at: %s", pathConfigfile)
	}

	configfile, err := filesystem.OpenFile(pathConfigfile)
	if err != nil {
		return nil, err
	}
	defer configfile.Close()

	var config Config
	err = json.NewDecoder(configfile).Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to decode config file: %s", err)
	}

	return &config, nil
}
