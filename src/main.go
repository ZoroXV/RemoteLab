package main

import (
    "log"
    "os"

    "remotelab/server"
    "remotelab/server/rest"
    "remotelab/upload"
)

var (
    DEFAULT_CONFIG_FILE string = "./config.json"
)

func main() {
    configFile := DEFAULT_CONFIG_FILE

    if len(os.Args) == 2 {
        configFile = os.Args[1]
    } else if len(os.Args) > 2 {
        log.Fatalf("[MAIN][ERR] Too much arguments.\n Use: %s [<config_file>]\n", os.Args[0])
    }

    upload.UploadInit()
    servers := server.CreateServers(configFile)

    for _, serv := range servers {
        if serv.Protocol == server.REST {
            rest.RunREST(serv)
        }
    }
}
