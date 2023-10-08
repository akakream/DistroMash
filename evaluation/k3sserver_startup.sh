curl -sfL https://get.k3s.io | sh - 

sudo chmod 777 /etc/rancher/k3s/k3s.yaml
export KUBECONFIG=/etc/rancher/k3s/k3s.yaml

sudo cp /var/lib/rancher/k3s/agent/etc/containerd/config.toml /var/lib/rancher/k3s/agent/etc/containerd/config.toml.tmpl

# Add the following to /var/lib/rancher/k3s/agent/etc/containerd/config.toml.tmpl
# [plugins."io.containerd.grpc.v1.cri".registry]
#   config_path = "/etc/containerd/certs.d"

sudo systemctl restart containerd
sudo systemctl restart k3s.service

git clone https://github.com/XenitAB/spegel.git

# Change spegel/charts/spegel/values.yaml with containerdSock: "/run/k3s/containerd/containerd.sock"

# Add Github token
export CR_PAT=YOUR_TOKEN
echo $CR_PAT | docker login ghcr.io -u USERNAME --password-stdin

helm upgrade --create-namespace --namespace spegel --install --version v0.0.12 spegel oci://ghcr.io/xenitab/helm-charts/spegel --values spegel/charts/spegel/values.yaml
