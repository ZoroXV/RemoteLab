echo Building remotelab:build

docker build -t remotelab:build . -f Dockerfile.build

docker container create --name extract remotelab:build
docker container cp extract:/app/remotelab ./remotelab
docker container rm -f extract

echo Building remotelab:latest

docker build --no-cache -t remotelab:latest .
rm ./remotelab
