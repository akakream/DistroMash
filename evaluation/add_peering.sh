#!/bin/bash

# Path to your JSON configuration file
config_file="~/.ipfs/config"

# Check if the configuration file exists
if [ -f "$config_file" ]; then
  # Check if IP address and PeerID were provided as script arguments
  if [ $# -ne 2 ]; then
    echo "Error: Please provide IP address and PeerID as script arguments."
    exit 1
  fi

  # Extract the IP address and PeerID from script arguments
  ip_address="$1"
  peer_id="$2"

  # Define the replacement data using script arguments
  replacement_data='[
    {
      "Addrs": [
        "/ip4/'"$ip_address"'/tcp/4001",
        "/ip4/'"$ip_address"'/udp/4001/quic-v1"
      ],
      "ID": "'"$peer_id"'"
    }
  ]'

  # Use jq to update the "Peers" key with the replacement data
  jq --argjson replacement "$replacement_data" '.Peering.Peers = $replacement' "$config_file" > temp_config.json

  # Replace the original config file with the updated one
  mv temp_config.json "$config_file"

  echo "Successfully updated the JSON configuration."
else
  echo "Error: JSON configuration file not found at $config_file"
fi

# Add IPFS bootstrap nodes with the provided values
ipfs bootstrap add /ip4/"$ip_address"/tcp/4001/p2p/"$peer_id"
ipfs bootstrap add /ip4/"$ip_address"/udp/4001/quic-v1/p2p/"$peer_id"

