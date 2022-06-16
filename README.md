# RemoteLab

## Authors
* Etienne CHANUDET (etienne.chanudet@epita.fr)
* Victor LE CORRE (victor.le-corre@epita.fr)
* Vincent MUSCEDERE (vincent.muscedere@epita.fr)

## Git Good Practices

### Coding Style

We use this [article](https://buzut.net/cours/versioning-avec-git/bien-nommer-ses-commits) as a reference.

### Git Forge

- Never merge **BROKEN** code on `master`
- Code on new branch for each feature
- Make **Pull Request** and get reviewed by the others

## Build & Run Project

### Build Docker Image
```sh
./build.sh
```

### Run Docker Image & Launch Server
```sh
docker run -p 8080:80 -v /dev:/dev -v /run/udev:/run/udev:ro --device-cgroup-rule='c 188:* rmw' remotelab:latest
```
