# DistroMash

DistroMash meshes your docker image distribution.

## TODO

- ADD LOGGER!!!

- Need some callback to get the CID from the Multiplatform2IPFS. It currently waits for the upload.
  Multiplatform2IPFS should get the order, do its logic and send the cid back.

- Do I need a database?? Everything is distributed through CRDT. Maybe to backup? In that case migrations? [https://github.com/golang-migrate/migrate#cli-usage](https://github.com/golang-migrate/migrate#cli-usage)

- Check `Makefile` and `Dockerfile` here in [https://github.com/create-go-app/fiber-go-template](https://github.com/create-go-app/fiber-go-template)
