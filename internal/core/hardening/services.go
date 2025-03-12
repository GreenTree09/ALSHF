package hardening

import (
    "fmt"
    "os/exec"
)

// Service represents a system service
type Service struct {
    Name string
}

// NewService creates a new service
func NewService(name string) Service {
    return Service{
        Name: name,
    }
}

// Start starts the service
func (s *Service) Start() error {
    cmd := exec.Command("systemctl", "start", s.Name)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to start service %s: %v", s.Name, err)
    }
    return nil
}

// Stop stops the service
func (s *Service) Stop() error {
    cmd := exec.Command("systemctl", "stop", s.Name)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to stop service %s: %v", s.Name, err)
    }
    return nil
}

// Status checks the status of the service
func (s *Service) Status() (string, error) {
    cmd := exec.Command("systemctl", "status", s.Name)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return "", fmt.Errorf("failed to get status of service %s: %v", s.Name, err)
    }
    return string(output), nil
}
