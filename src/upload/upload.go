package upload

import (
    "log"
    "os/exec"
)

func UploadInit() {
    log.Printf("[UPLOAD] Init upload module...\n")
    log.Printf("[UPLOAD] Load configuration...\n")
}

func UploadArduino(port string, cardType string, inputFile string) error {
    cmd := exec.Command(
        "arduino-cli",
        "upload",
        "-p", port,
        "--fqbn", cardType,
        "--input-file", inputFile)

    err := cmd.Run()
    if err != nil {
        log.Printf("[UPLOAD][ERR] Fail to upload '%s' on card '%s', port '%s'.\n\t%v\n", inputFile, cardType, port, err)
        return err
    }

    log.Printf("[UPLOAD] Upload successful\n")

    return nil
}
