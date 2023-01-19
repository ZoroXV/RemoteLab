package rest

type RestUploadFileHandler struct {}
type RestUploadHandler struct {}
type RestListControllersHandler struct {}
type RestServeCliScript struct {}

type restUploadRequest struct {
	SerialNumber  string `json:"serial_number"`
	StartAddress  string `json:"start_address"`
	Port          string `json:"port"`
	Fqbn          string `json:"fqbn"`
	FileName      string `json:"filename"`
}

type restResponse struct {
	Status        string `json:"status"`
	Message       string `json:"message"`
}

type restMicroControllerInfo struct {
	VendorName    string `json:"vendor_name"`
	ProductName   string `json:"product_name"`
	SerialNumber  string `json:"serial_number"`
	Port 		  string `json:"port"`
	Fqbn          []string `json:"fqbn"`
}

type restMicroControllersList struct {
	Status		  string `json:"status"`
	Message       string `json:"message"`
	Data		  []restMicroControllerInfo `json:"data"`
}
