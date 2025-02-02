package monitor

import (
	"os/exec"
)

func SetupTripwire() error {
	cmd := exec.Command("sudo", "tripwire", "--init")
	return cmd.Run()
}
