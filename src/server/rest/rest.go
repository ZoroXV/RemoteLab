package rest

import (
    "encoding/json"
    "errors"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"

    "remotelab/utils"
    "remotelab/server"
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
    if r.Method == "GET" {
	w.WriteHeader(http.StatusMethodNotAllowed)
    } else {
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
    }
}

func uploadCommand(port string, fqbn string, filename string) error {
    if !utils.FileExists(server.ROOT_FILE_PATH, filename) {
	return errors.New(fmt.Sprintf("File '%s' does not exists. Upload the file before running this command again.", filename))
    }

    // TODO: Call backend

    return nil
}

func handleCommandGET(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusMethodNotAllowed)
}

func handleCommandPOST(w http.ResponseWriter, r *http.Request) {
    var newReq restRequest
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
	    var err error

	    switch newReq.Command {
	    case "UPLOAD":
		err = uploadCommand(newReq.Port, newReq.Fqbn, newReq.FileName)

	    default:
		returnCode = http.StatusBadRequest
		resp = restResponse {
		    Status:	    "ERROR",
		    Message:    fmt.Sprintf("Invalid command '%s'.", newReq.Command),
		}
	    }

	    if err != nil {
		returnCode = http.StatusInternalServerError
		resp = restResponse {
		    Status: "ERROR",
		    Message: fmt.Sprintf("Something went wrong when executing the command '%s'. %v", newReq.Command, err),
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

func (this RestCommandHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
	handleCommandGET(w, r)
    } else if r.Method == "POST"{
	handleCommandPOST(w, r)
    } else {
	w.WriteHeader(http.StatusMethodNotAllowed)
    }
}
