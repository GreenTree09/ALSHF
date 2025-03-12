#!/bin/bash

# Load SELinux configuration from YAML file
config_file="configs/hardening/selinux.yaml"

# Check if yq is installed
if ! command -v yq &> /dev/null; then
    echo "yq command not found. Please install yq to continue."
    exit 1
fi

# Read the state and policy from the YAML file
state=$(yq e '.selinux.state' "$config_file")
policy=$(yq e '.selinux.policy' "$config_file")

# Check if state and policy are not empty
if [ -z "$state" ] || [ -z "$policy" ]; then
    echo "Failed to read SELinux configuration from $config_file"
    exit 1
fi

# Apply SELinux configuration
sudo setenforce "$state"
sudo semanage policy -m -t "$policy"

echo "SELinux configuration applied successfully."
