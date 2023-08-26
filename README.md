# DistroMash

DistroMash meshes your docker image distribution.

- It consists of 3 components: Controller, P2Pcomm and MultiPlatform2IPFS.
- Runs as a sidecar to IPFS daemon
- Leverages IPDR as docker registry to communicate with IPFS as image storage.

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

- Start ipfs daemon (installation instrcutions)
- Run DistroMash with `make run-all`

## TODO

- ADD LOGGER!!!

- Need some callback to get the CID from the Multiplatform2IPFS. It currently waits for the upload.
  Multiplatform2IPFS should get the order, do its logic and send the cid back.

- Do I need a database?? Everything is distributed through CRDT. Maybe to backup? In that case migrations? [https://github.com/golang-migrate/migrate#cli-usage](https://github.com/golang-migrate/migrate#cli-usage)

- Check `Makefile` and `Dockerfile` here in [https://github.com/create-go-app/fiber-go-template](https://github.com/create-go-app/fiber-go-template)
