# DistroMash

DistroMash meshes your docker image distribution.

- It consists of 3 components: Controller, P2Pcomm and MultiPlatform2IPFS.
- Runs as a sidecar to IPFS daemon
- Leverages IPDR as docker registry to communicate with IPFS as image storage.

## Install

- `git submodule init && git submodule update`
- go version 1.20.5: [https://go.dev/doc/install](https://go.dev/doc/install)
- IPFS Kubo 0.20.0 or above [https://docs.ipfs.tech/install/command-line/#system-requirements](https://docs.ipfs.tech/install/command-line/#system-requirements)
- Docker

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
  - Run that node as shown below to create a multi-address for it.
  - The multi-address will be printed to standard output. Grab it and continue below.
  - DistroMash uses a file for bootstrapping: `data/peerstore`.
  - There is a `P2PComm/data/peerstore.template` file showing how the bootstrap entry look like.
  - Copy the `P2PComm/data/peerstore.template` file to `P2PComm/data/peerstore`.
  - Add the multi-address of the bootstrap node to every node's `P2PComm/data/peerstore`.

## Run

- Start ipfs daemon with `ipfs daemon`
- Run DistroMash with `make run-all`
- To see the web application, go to `localhost:3000`
- To see the API, go to `localhost:3000/swagger`

## How to register a strategy?

There are two types of strategies: `percentage` and `target`:

- `percentage` strategy replicates the Docker image in the given percentage of the edge environment. An example payload for `percentage` strategy looks like this:

  ```{
        "execute": true,
        "nametag": "busybox:1.35.0",
        "percentage": 50,
        "target": "doesntmatter",
        "type": "percentage"
    }
  ```

  This strategy replicates busybox:1.35.0 image to the 50% of the nodes in the edge environment. It starts the execution rightaway.

- `target` strategy replicates the Docker image to a specific edge node (target) in the edge environment. An example payload for `target` strategy looks like this:

  ```{
      "execute": true,
      "nametag": "busybox:1.35.0",
      "percentage": 0,
      "target": "QmUdSkPo6Q3U3PmOJZM4Pk2Y5z5pJ8fyj433HQ4tUQbkX7",
      "type": "target"
  }
  ```

  This strategy replicates busybox:1.35.0 image to the target node in the edge environment. It starts the execution rightaway.

## Web UI

DistroMash also offers a web UI to explore the registered strategies, its distributed data store and the peers in the peer-to-peer network.

- To check out the web UI, SSH tunnel to port 3000 (default) of one of the peers that runs DistroMash.

```
ssh -i ~/.ssh/gcp-instance-1 -L 3000:localhost:3000 username@IP
```

## TODO

- When strategy is removed, `docker rmi` images from nodes
