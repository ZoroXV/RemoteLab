package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

func (this *Server) runServer(handler http.Handler) {
	hostname, _ := os.Hostname()

	this.Running = true
	log.Printf("[%s][HTTP] Listening on %s:%s ...\n", this.Protocol, hostname, this.Port)

	if err := http.ListenAndServe(fmt.Sprintf(":%v", this.Port), handler); err != nil {
		this.Running = false
		log.Printf("[%s][HTTP][ERR] Cannot launch server on %s:%s.\n\t%v\n", this.Protocol, hostname, this.Port, err)
	}
}

func CreateServers(configFile string) []Server {
	return ParseConfigFile(configFile)
}

func (this *Server) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	s := http.NewServeMux()

	for _, h := range this.Handlers {
		s.Handle(h.path, this.Logger(h.handler))
	}

	this.runServer(s)
}

func (this *Server) AddHandler(path string, handl http.Handler) {
	if !this.Running {
		this.Handlers = append(this.Handlers, handler{path, handl})
		log.Printf("[%s][HANDLER] Add handler for '%s'.", this.Protocol, path)
	} else {
		log.Printf("[%s][HANDLER][ERR] Cannot add handler for '%s'. Server is running.", this.Protocol, path)
	}
}
