# ALSHF
## (Automated Linux Security Hardening Framework)

![License](https://img.shields.io/github/license/GreenTree09/ALSHF)
![Go Version](https://img.shields.io/github/go-mod/go-version/GreenTree09/ALSHF)
![Build Status](https://img.shields.io/github/workflow/status/GreenTree09/ALSHF/CI)

ALSHF is a comprehensive security automation tool designed to enhance Linux system security through automated hardening, continuous monitoring, and real-time threat detection.

## 🚀 Features

### System Hardening
- Automated system hardening based on security best practices
- Firewall configuration and management (UFW/iptables)
- SELinux/AppArmor configuration
- Service hardening and management
- Secure configuration of critical system files

### Security Monitoring
- File integrity monitoring (AIDE/Tripwire integration)
- Real-time intrusion detection
- System anomaly detection
- Comprehensive logging and auditing

### Vulnerability Management
- Automated vulnerability scanning
- Integration with Lynis and OpenVAS
- Scheduled security audits
- Automated patch management

### Alert System
- Real-time security alerts
- Multiple notification channels (Email, Slack, Telegram)
- Customizable alert rules
- Detailed security reports

## 🛠 Installation

```bash
# Clone the repository
git clone https://github.com/GreenTree09/ALSHF.git

# Change into the directory
cd ALSHF

# Build the project
make build

# Install dependencies
make deps
```

## 📋 Prerequisites

- Go 1.21 or higher
- Linux-based operating system (Ubuntu 20.04+, Debian 11+, CentOS 8+)
- Root/sudo access
- Required system tools:
  - AIDE or Tripwire
  - UFW or iptables
  - Lynis
  - OpenVAS (optional)

## 🚦 Quick Start

```bash
# Install ALSHF
sudo ./install.sh

# Run initial system hardening
sudo alshf harden --config /etc/alshf/config.yaml

# Start monitoring service
sudo systemctl start alshfd

# Check status
alshf status
```

## 📖 Documentation

- [Installation Guide](docs/installation.md)
- [Configuration Guide](docs/configuration.md)
- [Hardening Features](docs/hardening.md)
- [Monitoring Setup](docs/monitoring.md)
- [Alert Configuration](docs/alerts.md)
- [API Documentation](docs/api.md)
- [Contributing Guide](CONTRIBUTING.md)

## ⚙️ Configuration

Basic configuration example:

```yaml
hardening:
  firewall:
    enabled: true
    default_policy: DROP
  monitoring:
    aide_check_interval: 3600
    alert_channels:
      - type: telegram
      - type: email
```

## 🤝 Contributing

Contributions are welcome! Please read our [Contributing Guide](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.

## 🔒 Security

For security issues, please read our [Security Policy](SECURITY.md) and report any vulnerabilities responsibly.

## ✨ Acknowledgments

- Linux Security Community
- Go Community
- All Contributors

## 📊 Project Status

Current Version: 0.1.0 (Development)  
Last Updated: 2025-01-20 22:54:10 UTC

This project is under active development. Features are being added regularly, and the API might change until version 1.0.0.

## 📞 Contact

Project Link: [https://github.com/GreenTree09/ALSHF](https://github.com/GreenTree09/ALSHF)

---
Created with ❤️ by GreenTree09
