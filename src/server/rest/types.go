package rest

type RestUploadFileHandler struct {}
type RestCommandHandler struct {}

type restRequest struct {
    Command	string `json:"command"`
    Port	string `json:"port"`
    Fqbn	string `json:"fqbn"`
    FileName	string `json:"filename"`
}

type restResponse struct {
    Status	string `json:"status"`
    Message	string `json:"message"`
}

