package main

import (
	"fmt"
	"internal/core/hardening"
	"internal/core/monitor"
	"log"
	"net/http"
)

func main() {
	// Setup Firewall
	if err := hardening.SetupFirewall("configs/hardening/firewall.yaml"); err != nil {
		log.Fatalf("Error setting up firewall: %v", err)
	}

	// Configure SELinux
	if err := hardening.ConfigureSELinux("configs/hardening/selinux.yaml"); err != nil {
		log.Fatalf("Error configuring SELinux: %v", err)
	}

	// Setup AIDE
	if err := monitor.SetupAIDE("configs/monitoring/aide.yaml"); err != nil {
		log.Fatalf("Error setting up AIDE: %v", err)
	}

	// Check Integrity
	if err := monitor.CheckIntegrity(); err != nil {
		log.Fatalf("Error checking integrity: %v", err)
	}

	// Setup Tripwire
	if err := monitor.SetupTripwire(); err != nil {
		log.Fatalf("Error setting up Tripwire: %v", err)
	}

	// Start HTTP server for monitoring
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ALSHF Daemon is running")
	})

	log.Println("Starting ALSHF Daemon Service...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
