# NODE_TOKEN comes from /var/lib/rancher/k3s/server/node-token on your server
curl -sfL https://get.k3s.io | K3S_URL=https://myserver:6443 K3S_TOKEN=mynodetoken sh -

sudo cp /var/lib/rancher/k3s/agent/etc/containerd/config.toml /var/lib/rancher/k3s/agent/etc/containerd/config.toml.tmpl

# Add the following to /var/lib/rancher/k3s/agent/etc/containerd/config.toml.tmpl
# [plugins."io.containerd.grpc.v1.cri".registry]
#   config_path = "/etc/containerd/certs.d"

sudo systemctl restart containerd
sudo systemctl restart k3s-agent
