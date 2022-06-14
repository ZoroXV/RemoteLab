package rest

type RestUploadFileHandler struct {}
type RestUploadHandler struct {}

type restUploadRequest struct {
	Port          string `json:"port"`
	Fqbn          string `json:"fqbn"`
	FileName      string `json:"filename"`
}

type restResponse struct {
	Status        string `json:"status"`
	Message       string `json:"message"`
}

