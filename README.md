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

### Run Docker Image in RELEASE mode
```sh
docker compose up release
```

### Raspberry Pi Daemon

We use a simple daemon to launch our server at the start of the Raspberry Pi.
Our daemon will restart the server if a error occur.
Our daemon execute a simple bash script that run our Docker Image.
