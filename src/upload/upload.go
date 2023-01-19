package upload

import (
    "log"
    "os/exec"
    "strconv"
    "strings"
)

func UploadInit() {
    log.Printf("[UPLOAD] Init upload module...\n")
    log.Printf("[UPLOAD] Load configuration...\n")
}

func Upload(serialNumber string, startAddress string, port string, cardType string, inputFile string) error {
    if cardType != "" {
        return uploadArduino(port, cardType, inputFile)
    } else {
        return uploadSTM(serialNumber, startAddress, inputFile)
    }
}

func uploadArduino(port string, cardType string, inputFile string) error {
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

func uploadSTM(serial string, startAddress string, inputFile string) error {
    startAddressCleaned := strings.Replace(startAddress, "0x", "", -1)
    _, err := strconv.ParseInt(startAddressCleaned, 16, 64)

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
        log.Printf("[UPLOAD][STM32][ERR] Fail to upload '%s' on card 'STM32', serial '%s'.\n\t%v\n", inputFile, serial, err)
        return err
    }

    log.Printf("[UPLOAD][STM32] Upload successful\n")

    return nil
}