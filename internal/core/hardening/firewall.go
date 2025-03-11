package hardening

import (
	"io/ioutil"
	"os/exec"
	"gopkg.in/yaml.v2"
)

type FirewallConfig struct {
	Rules []struct {
		Allow string `yaml:"allow"`
	} `yaml:"rules"`
}

func SetupFirewall(configPath string) error {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	var config FirewallConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	cmd := exec.Command("sudo", "ufw", "enable")
	if err := cmd.Run(); err != nil {
		return err
	}

	for _, rule := range config.Rules {
		cmd := exec.Command("sudo", "ufw", "allow", rule.Allow)
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}
