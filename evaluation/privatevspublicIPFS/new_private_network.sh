echo -e "/key/swarm/psk/1.0.0/\n/base16/\n`tr -dc 'a-f0-9' < /dev/urandom | head -c64`" > ~/.ipfs/swarm.key

scp -i ~/.ssh/gcp-instance-1 akakream@34.107.53.120:~/.ipfs/swarm.key .

scp  -i ~/.ssh/gcp-instance-1 swarm.key akakream@35.242.230.77:~/.ipfs/swarm.key
scp  -i ~/.ssh/gcp-instance-1 swarm.key akakream@35.198.66.248:~/.ipfs/swarm.key
scp  -i ~/.ssh/gcp-instance-1 swarm.key akakream@34.159.130.93:~/.ipfs/swarm.key


ipfs config show | grep "PeerID"
# At every node
ipfs bootstrap add /ip4/10.156.0.19/tcp/4001/ipfs/12D3KooWAZCEcytNJGRAuWBv2xJqAZfMmF3S8BRuwtoqW2CzZaJw

export LIBP2P_FORCE_PNET=1 
ipfs daemon
