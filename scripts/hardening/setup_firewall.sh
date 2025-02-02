#!/bin/bash
# Load firewall rules from YAML configuration
config_file="configs/hardening/firewall.yaml"

# Enable UFW
sudo ufw enable

# Parse YAML and apply rules
rules=$(yq e '.firewall.rules[].allow' $config_file)
for rule in $rules; do
    sudo ufw allow $rule
done