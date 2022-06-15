package main

import (
	"log"
	"os"
	"sync"

	"remotelab/server"
	"remotelab/server/rest"
	"remotelab/upload"
)

var (
	DEFAULT_CONFIG_FILE string = "./config.json"
)

func main() {
	configFile := DEFAULT_CONFIG_FILE

	if len(os.Args) == 1 {
		server.CreateDefaultConfig(configFile)
	} else if len(os.Args) == 2 {
		configFile = os.Args[1]
	} else if len(os.Args) > 2 {
		log.Fatalf("[MAIN][ERR] Too much arguments.\n Use: %s [<config_file>]\n", os.Args[0])
	}

	wg := new(sync.WaitGroup)
	upload.UploadInit()
	servers := server.CreateServers(configFile)

	for _, serv := range servers {
		if serv.Protocol == server.REST {
			wg.Add(1)
			go rest.RunREST(serv, wg)
		}
	}

	wg.Wait()
}
