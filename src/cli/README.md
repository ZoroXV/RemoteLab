# CLI tool description

### Download
To get the CLI tool, the easiest way is to download it from the running server:
```sh
wget <raspberrypi_ip>/download/remotelab.py
chmod +x remotelab.py
```

### Usage
```sh
./remotelab.py [-h] {flash,upload,list}
```

#### Flash
```sh
./remotelab.py flash [-h] -a ADDRESS -f FQBN -p PORT filename

Arguments:
  filename              the name of the file to flash on the microcontroller
  -h, --help            show this help message and exit
  -a ADDRESS, --address ADDRESS
                        the ip address of the server
  -f FQBN, --fqbn FQBN  the type of the card, following the names of the `arduino-cli` (ex: "arduino:avr:uno")
  -p PORT, --port PORT  the port on which the card is linked (ex: "/dev/ttyUSB0")
```

#### Upload files
```sh
./remotelab.py upload [-h] -a ADDRESS filepath [filepath ...]

Arguments:
  filepath              the full path of the file(s) to upload on the server
  -h, --help            show this help message and exit
  -a ADDRESS, --address ADDRESS
                        the ip address of the server
```

#### List microcontrollers
```sh
./remotelab.py list [-h] -a ADDRESS

Arguments:
  -h, --help            show this help message and exit
  -a ADDRESS, --address ADDRESS
                        the ip address of the server
```