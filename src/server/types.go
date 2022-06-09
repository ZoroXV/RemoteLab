package server

import (
    "net/http"
)

type Vhosts struct {
    Vhosts	[]Config `json:"vhosts"`
}

type Config struct {
    Protocol	string `json:"protocol"`
    Port	string `json:"port"`
}

type Protocol string

const (
    REST Protocol = "REST"
)

type handler struct {
    path	    string
    handler	    http.Handler
}

type Server struct {
    Protocol	    Protocol
    Port	    string
    Handlers	    []handler
    Running	    bool
}
