package rest

import (
    "encoding/json"
    "errors"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "sync"

    "remotelab/utils"
    "remotelab/server"
    "remotelab/upload"
    "remotelab/monitor"
)

func fileSave(r *http.Request) error {
    err := r.ParseMultipartForm(32 << 20)
    if err != nil {
        return err
    }

    filename := r.Form.Get("name")
    file, _, err := r.FormFile("file")
    if err != nil {
        return err
    }

    defer file.Close()

    return utils.SaveFile(server.ROOT_FILE_PATH, filename, file)
}

func (this RestUploadFileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        resp := restResponse {
            Status:	"OK",
            Message:    "",
        }
        returnCode := http.StatusOK

        err := fileSave(r)

        if err != nil {
            returnCode = http.StatusInternalServerError
            resp = restResponse {
                Status:	    "ERROR",
                Message:    fmt.Sprintf("Command upload fail. %v", err),
            }
        }

        jsonResp, err := json.Marshal(resp)
        if err != nil {
            log.Fatalf("[REST][HANDLE_REQUEST][ERR] Fail to encode json.\n\t%v\n", err)
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(returnCode)
        w.Write(jsonResp)
    } else {
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

func uploadCommand(port string, fqbn string, filename string) error {
    if !utils.FileExists(server.ROOT_FILE_PATH, filename) {
        return errors.New(fmt.Sprintf("File '%s' does not exists. Upload the file before running this command again.", filename))
    }

    return upload.UploadArduino(port, fqbn, utils.GetFullPath(server.ROOT_FILE_PATH, filename))
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
    var newReq restUploadRequest
    resp := restResponse {
        Status:	    "OK",
        Message:    "",
    }
    returnCode := http.StatusOK

    reqBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Printf("[REST][HANDLE_REQUEST][ERR] Fail to read the body of the request.\n\t%v\n", err)

        returnCode = http.StatusBadRequest
        resp = restResponse {
            Status:	"ERROR",
            Message:	fmt.Sprintf("Fail to read the body of the request. %v", err),
        }
    } else {
        err = json.Unmarshal(reqBody, &newReq)
        if err != nil {
            log.Printf("[REST][HANDLE_REQUEST][ERR] Fail to decode json.\n\t%v\n", err)

            returnCode = http.StatusBadRequest
            resp = restResponse {
                Status:	    "ERROR",
                Message:    fmt.Sprintf("Fail to decode json. %v", err),
            }
        } else {
            err := uploadCommand(newReq.Port, newReq.Fqbn, newReq.FileName)

            if err != nil {
                returnCode = http.StatusInternalServerError
                resp = restResponse {
                    Status: "ERROR",
                    Message: fmt.Sprintf("Something went wrong when executing the upload command. %v", err),
                }
            }
        }
    }

    jsonResp, err := json.Marshal(resp)
    if err != nil {
        log.Fatalf("[REST][HANDLE_REQUEST][ERR] Fail to encode json.\n\t%v\n", err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(returnCode)
    w.Write(jsonResp)
}

func (this RestUploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        handleUpload(w, r)
    } else {
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

func (this RestListControllersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        microContInfos := monitor.ListMicrocontrollers()

        var resp []restMicroControllerInfo

        for _, controller := range microContInfos {
            fqbn, err := upload.GetFqbn(controller.VendorID, controller.ProductID)

            if err != nil {
                fqbn = fmt.Sprint(err)
            }

            resp := append(resp, restMicroControllerInfo{
                VendorName: controller.VendorName,
	            ProductName: controller.ProductName,
	            Port: controller.Port,
	            Fqbn: fqbn,
            })
        }

        jsonResp, err := json.Marshal(resp)

        if err != nil {
            log.Fatalf("[REST][HANDLE_REQUEST][ERR] Fail to encode json.\n\t%v\n", err)
        }
    
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(jsonResp)
    } else {
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

func RunREST(serv server.Server, wg *sync.WaitGroup) {
    restUploadFileHandler := RestUploadFileHandler{}
    serv.AddHandler("/uploadfile", restUploadFileHandler)

    restUploadHandler := RestUploadHandler{}
    serv.AddHandler("/command/upload", restUploadHandler)

    restListControllersHandler := RestListControllersHandler{}
    serv.AddHandler("/command/list_controllers", restListControllersHandler)

    serv.Run(wg)
}
