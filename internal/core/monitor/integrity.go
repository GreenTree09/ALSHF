package monitor

import (
	"os/exec"
)

func CheckIntegrity() error {
	cmd := exec.Command("sudo", "aide", "--check")
	return cmd.Run()
}
