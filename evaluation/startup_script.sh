for pkg in docker.io docker-doc docker-compose podman-docker containerd runc; do sudo apt-get remove $pkg; done

# Add Docker's official GPG key:
sudo apt-get update
sudo apt-get install ca-certificates curl gnupg
sudo install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
sudo chmod a+r /etc/apt/keyrings/docker.gpg

# Add the repository to Apt sources:
echo \
  "deb [arch="$(dpkg --print-architecture)" signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  "$(. /etc/os-release && echo "$VERSION_CODENAME")" stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update

sudo apt-get -y install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

sudo chmod 777 /var/run/docker.sock

curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3

chmod 700 get_helm.sh

./get_helm.sh

wget https://go.dev/dl/go1.20.5.linux-amd64.tar.gz

sudo tar -C /usr/local -xzf go1.20.5.linux-amd64.tar.gz

export PATH=$PATH:/usr/local/go/bin

wget https://dist.ipfs.tech/kubo/v0.22.0/kubo_v0.22.0_linux-amd64.tar.gz

tar -xvzf kubo_v0.22.0_linux-amd64.tar.gz

cd kubo

sudo bash install.sh

ipfs init

sudo apt install make

git clone https://github.com/akakream/DistroMash.git

cd ..

git clone https://github.com/akakream/DistroMash.git

echo '127.0.0.1 docker.local' | sudo tee -a /etc/hosts
echo '::1       docker.local' | sudo tee -a /etc/hosts

cd DistroMash

git submodule init && git submodule update

make build-all

cd ..

sudo apt install net-tools
