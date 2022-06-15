package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"remotelab/utils"
)

func ParseConfigFile(configFile string) []Server {
	var data Vhosts
	var result []Server

	jsonFile, err := os.Open(configFile)
	if err != nil {
		log.Fatalf("[CONFIG][ERR] Cannot open config file '%s'.\n\t%v\n", configFile, err)
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("[CONFIG][ERR] Cannot read config file '%s'.\n\t%v\n", configFile, err)
	}


	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		log.Fatalf("[CONFIG][ERR] Cannot parse config file '%s'.\n\t%v\n", configFile, err)
	}

	for _, e := range data.Vhosts {
		server := Server {
			Protocol:	Protocol(e.Protocol),
			Port:   	e.Port,
			Running:	false,
			Handlers:	nil,
		}

		result = append(result, server)
	}

	return result
}

func CreateDefaultConfig(filename string) error {
	rootPath, _ := os.Getwd()

	if utils.FileExists(rootPath, filename) {
		return nil
	}

	outputFile, err := os.OpenFile(filename, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	defer outputFile.Close()

	_, err = outputFile.WriteString(`{"vhosts":[{"protocol":"REST","port":"8080"}]}`)
	if err != nil {
		return err
	}

	return nil
}
