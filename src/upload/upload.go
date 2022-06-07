package upload

import (
	"fmt"
	"log"
	"os/exec"
)

func UploadInit() {
	fmt.Println("Init Upload Module")
	fmt.Println("Load Configuration...")
}

func UploadArduino(port string, cardType string, inputFile string) {
	cmd := "arduino-cli upload -p " + port + " --fqbn " + cardType + " --input-file " + inputFile

	out, err := exec.Command(cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Upload Sucessful")
	output := string(out[:])
	fmt.Println(output)
}
