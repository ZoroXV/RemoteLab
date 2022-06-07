package upload

import (
	"fmt"
)

func UploadInit() {
	fmt.Println("Init Upload Module")
	fmt.Println("Load Configuration...")
}

func UploadArduino() {
	fmt.Println("arduino-cli upload -p <port> --fqbn <typeDeCarte> --input-file <binaire>")
}
