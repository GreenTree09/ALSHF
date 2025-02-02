package main

import (
	"fmt"
	"internal/core/hardening"
	"internal/core/monitor"
)

func main() {
	// Setup Firewall
	if err := hardening.SetupFirewall("configs/hardening/firewall.yaml"); err != nil {
		fmt.Println("Error setting up firewall:", err)
	}

	// Configure SELinux
	if err := hardening.ConfigureSELinux("configs/hardening/selinux.yaml"); err != nil {
		fmt.Println("Error configuring SELinux:", err)
	}

	// Setup AIDE
	if err := monitor.SetupAIDE("configs/monitoring/aide.yaml"); err != nil {
		fmt.Println("Error setting up AIDE:", err)
	}

	// Check Integrity
	if err := monitor.CheckIntegrity(); err != nil {
		fmt.Println("Error checking integrity:", err)
	}

	// Setup Tripwire
	if err := monitor.SetupTripwire(); err != nil {
		fmt.Println("Error setting up Tripwire:", err)
	}
}
