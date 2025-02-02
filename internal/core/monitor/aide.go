package monitor

import (
	"os/exec"

	"gopkg.in/yaml.v2"
)

type AIDEConfig struct {
	Config   string `yaml:"config"`
	Database string `yaml:"database"`
}

func SetupAIDE(configPath string) error {
	data, err := io.ReadFile(configPath)
	if err != nil {
		return err
	}

	var config AIDEConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	cmd := exec.Command("sudo", "aide", "--init", "--config", config.Config, "--database", config.Database)
	return cmd.Run()
}
