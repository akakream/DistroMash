# DistroMash

DistroMash meshes your docker image distribution.

- It consists of 3 components: Controller, P2Pcomm and MultiPlatform2IPFS.
- Runs as a sidecar to IPFS daemon
- Leverages IPDR as docker registry to communicate with IPFS as image storage.

## Install

- `git submodule init && git submodule update`
- go version 1.20.5: [https://go.dev/doc/install](https://go.dev/doc/install)
- IPFS Kubo 0.20.0 or above [https://docs.ipfs.tech/install/command-line/#system-requirements](https://docs.ipfs.tech/install/command-line/#system-requirements)

## Build

- Build the micro-services with `make build-all`
- Add `docker.local` to `/etc/hosts`:

  ```hosts
  echo '127.0.0.1 docker.local' | sudo tee -a /etc/hosts
  echo '::1       docker.local' | sudo tee -a /etc/hosts
  ```

  - Flush local DNS cache:

    - on macOS:

      ```bash
      dscacheutil -flushcache; sudo killall -HUP mDNSResponder
      ```

    - on Ubuntu 18+:

      ```bash
      sudo systemd-resolve --flush-caches
      ```

## Bootstrap

- Choose a node or nodes from your cluster and boostrap it.
  - DistroMash uses a file for bootstrapping: `data/peerstore`.
  - There is a `data/peerstore.template` file showing how the bootstrap entry look like.
  - Copy `data/peerstore.template` file to `data/peerstore`.
  - Add the multi-address of the bootstrap node to every node's `data/peerstore`.

## Run

- Start ipfs daemon with `ipfs daemon`
- Run DistroMash with `make run-all`
