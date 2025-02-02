#!/bin/bash
# Load SELinux configuration from YAML file
config_file="configs/hardening/selinux.yaml"

state=$(yq e '.selinux.state' $config_file)
policy=$(yq e '.selinux.policy' $config_file)

# Apply SELinux configuration
sudo setenforce $state
sudo semanage policy -m -t $policy