#!/bin/bash
# Load AIDE configuration from YAML file
config_file="configs/monitoring/aide.yaml"

config=$(yq e '.aide.config' $config_file)
database=$(yq e '.aide.database' $config_file)

# Initialize AIDE
sudo aide --init --config $config --database $database