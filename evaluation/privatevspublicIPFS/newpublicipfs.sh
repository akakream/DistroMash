ssh -i ~/.ssh/gcp-instance-1 akakream@35.242.230.77 "sudo rm -r ~/.ipfs"
ssh -i ~/.ssh/gcp-instance-1 akakream@35.198.66.248 "sudo rm -r ~/.ipfs"
ssh -i ~/.ssh/gcp-instance-1 akakream@34.159.130.93 "sudo rm -r ~/.ipfs"

ssh -i ~/.ssh/gcp-instance-1 akakream@35.242.230.77 "ipfs init"
ssh -i ~/.ssh/gcp-instance-1 akakream@35.198.66.248 "ipfs init"
ssh -i ~/.ssh/gcp-instance-1 akakream@34.159.130.93 "ipfs init"

ssh -i ~/.ssh/gcp-instance-1 akakream@35.242.230.77 "ipfs config Discovery.MDNS.Enabled --bool false"
ssh -i ~/.ssh/gcp-instance-1 akakream@35.198.66.248 "ipfs config Discovery.MDNS.Enabled --bool false"
ssh -i ~/.ssh/gcp-instance-1 akakream@34.159.130.93 "ipfs config Discovery.MDNS.Enabled --bool false"
