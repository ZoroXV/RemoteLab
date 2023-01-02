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

## Install RemoteLab

```sh
chmod +x install.sh # If the script does not have the right to be executed
./install.sh
```

Running this command install RemoteLab to `$HOME/.remotelab` and enable a service called `remotelab.service` which run and restart automatically the server.

## Build & Run Project
```sh
cd <remotelab_location>
```

### Build
```sh
docker compose build
```

### Run
```sh
docker compose up release
```
