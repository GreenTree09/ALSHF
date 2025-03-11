package hardening

import (
	"os/exec"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type SELinuxConfig struct {
	State  string `yaml:"state"`
	Policy string `yaml:"policy"`
}

func ConfigureSELinux(configPath string) error {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	var config SELinuxConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	cmd := exec.Command("sudo", "setenforce", config.State)
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("sudo", "semanage", "policy", "-m", "-t", config.Policy)
	return cmd.Run()
}
