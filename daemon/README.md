### Raspberry Pi Daemon

We use a simple daemon using `systemctl` to launch our server at the start of the Raspberry Pi.
The daemon will restart the server if an error occur.
It daemon execute a simple bash script that run the Remotelab Docker Image.