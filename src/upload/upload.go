package upload

import (
    "log"
    "os/exec"
    "strconv"
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
        log.Printf("[UPLOAD][ARDUINO][ERR] Fail to upload '%s' on card '%s', port '%s'.\n\t%v\n", inputFile, cardType, port, err)
        return err
    }

    log.Printf("[UPLOAD][ARDUINO] Upload successful\n")

    return nil
}

func UploadSTM(port string, serial string, inputFile string, startAddress string) error {
    _, err := strconv.ParseInt(startAddress, 16, 64)

    if err != nil {
        log.Printf("[UPLOAD][STM32][ERR] Invalid start address '%s'.\n\t%v\n", startAddress, err)
        return err
    }

    cmd := exec.Command(
        "st-flash",
        "--serial", serial,
        "write", inputFile,
        startAddress)

    err = cmd.Run()
    if err != nil {
        log.Printf("[UPLOAD][STM32][ERR] Fail to upload '%s' on card 'STM32', serial '%s', port '%s'.\n\t%v\n", inputFile, serial, port, err)
        return err
    }

    log.Printf("[UPLOAD][STM32] Upload successful\n")

    return nil
}