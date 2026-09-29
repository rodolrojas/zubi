package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"rodolrojas.com/zubi/internal/common/structs"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func LoadConfig() (*structs.BaseServerConfig, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(filepath.Join(currentDir, "config.yml"))
	if err != nil {
		return nil, err
	}
	logrus.Infof("Loaded config.yml successfully")

	config := &structs.BaseConfig{}

	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, err
	}
	// DebugConfig(&config.Server)
	
	return &config.Server, nil
}

func DebugConfig(config *structs.BaseServerConfig) {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		logrus.Errorf("Error marshalling config for debug: %v", err)
		return
	}
	logrus.Debugf("Current configuration:\n%s", string(data))
}