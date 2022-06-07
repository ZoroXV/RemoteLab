package main

import (
	"remotelab/upload"
)

func main() {
	upload.UploadInit()

	upload.UploadArduino("COM3", "arduino:avr:uno",
		"/home/victor/Documents/RemoteLab/controllers/arduino/blink_1/blink_1.ino.with_bootloader.bin")
}
