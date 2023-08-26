# DistroMash

DistroMash meshes your docker image distribution.

- It consists of 3 components: Controller, P2Pcomm and MultiPlatform2IPFS.
- Runs as a sidecar to IPFS daemon
- Leverages IPDR as docker registry to communicate with IPFS as image storage.

## Install

- `git submodule init && git submodule update`
- go version 1.20.5 or above: [https://go.dev/doc/install](https://go.dev/doc/install)
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

## Run

- Start ipfs daemon with `ipfs daemon`
- Run DistroMash with `make run-all`
