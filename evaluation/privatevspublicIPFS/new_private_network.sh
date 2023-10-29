ssh -i ~/.ssh/gcp-instance-1 akakream@34.107.53.120 "sudo rm -r ~/.ipfs"
ssh -i ~/.ssh/gcp-instance-1 akakream@35.242.230.77 "sudo rm -r ~/.ipfs"
ssh -i ~/.ssh/gcp-instance-1 akakream@35.198.66.248 "sudo rm -r ~/.ipfs"
ssh -i ~/.ssh/gcp-instance-1 akakream@34.159.130.93 "sudo rm -r ~/.ipfs"

ssh -i ~/.ssh/gcp-instance-1 akakream@34.107.53.120 "ipfs init"
ssh -i ~/.ssh/gcp-instance-1 akakream@35.242.230.77 "ipfs init"
ssh -i ~/.ssh/gcp-instance-1 akakream@35.198.66.248 "ipfs init"
ssh -i ~/.ssh/gcp-instance-1 akakream@34.159.130.93 "ipfs init"

ssh -i ~/.ssh/gcp-instance-1 akakream@34.107.53.120 'echo -e "/key/swarm/psk/1.0.0/\n/base16/\n`tr -dc 'a-f0-9' < /dev/urandom | head -c64`" > ~/.ipfs/swarm.key'

scp -i ~/.ssh/gcp-instance-1 akakream@34.107.53.120:~/.ipfs/swarm.key .

scp -i ~/.ssh/gcp-instance-1 swarm.key akakream@35.242.230.77:~/.ipfs/swarm.key
scp -i ~/.ssh/gcp-instance-1 swarm.key akakream@35.198.66.248:~/.ipfs/swarm.key
scp -i ~/.ssh/gcp-instance-1 swarm.key akakream@34.159.130.93:~/.ipfs/swarm.key

ssh -i ~/.ssh/gcp-instance-1 akakream@34.107.53.120 "ipfs bootstrap rm --all"
ssh -i ~/.ssh/gcp-instance-1 akakream@35.242.230.77 "ipfs bootstrap rm --all"
ssh -i ~/.ssh/gcp-instance-1 akakream@35.198.66.248 "ipfs bootstrap rm --all"
ssh -i ~/.ssh/gcp-instance-1 akakream@34.159.130.93 "ipfs bootstrap rm --all"

#ssh -i ~/.ssh/gcp-instance-1 akakream@34.107.53.120 "ipfs bootstrap add /ip4/10.156.0.19/tcp/4001/ipfs/12D3KooWAWRsxySvZndVjBn5aTCQXv67SgR6YE41tq9877mfiYxg"
#ssh -i ~/.ssh/gcp-instance-1 akakream@35.242.230.77 "ipfs bootstrap add /ip4/10.156.0.19/tcp/4001/ipfs/12D3KooWAWRsxySvZndVjBn5aTCQXv67SgR6YE41tq9877mfiYxg"
#ssh -i ~/.ssh/gcp-instance-1 akakream@35.198.66.248 "ipfs bootstrap add /ip4/10.156.0.19/tcp/4001/ipfs/12D3KooWAWRsxySvZndVjBn5aTCQXv67SgR6YE41tq9877mfiYxg"
#ssh -i ~/.ssh/gcp-instance-1 akakream@34.159.130.93 "ipfs bootstrap add /ip4/10.156.0.19/tcp/4001/ipfs/12D3KooWAWRsxySvZndVjBn5aTCQXv67SgR6YE41tq9877mfiYxg"

rm swarm.key

# export LIBP2P_FORCE_PNET=1 
# ipfs daemon

