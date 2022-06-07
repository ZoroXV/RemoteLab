package main

import (
	"fmt"

	"remotelab/upload"
)

func main() {
	message := upload.UploadInit()
	fmt.Println(message)
}
