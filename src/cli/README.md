# CLI tool description

### Usage
```sh
./remotelab [-h] {flash,upload,list}
```

#### Flash
```sh
./remotelab flash [-h] -a ADDRESS -f FQBN -p PORT filename

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
./remotelab upload [-h] -a ADDRESS filepath [filepath ...]

Arguments:
  filepath              the full path of the file(s) to upload on the server
  -h, --help            show this help message and exit
  -a ADDRESS, --address ADDRESS
                        the ip address of the server
```

#### List microcontrollers
```sh
./remotelab list [-h] -a ADDRESS

Arguments:
  -h, --help            show this help message and exit
  -a ADDRESS, --address ADDRESS
                        the ip address of the server
```