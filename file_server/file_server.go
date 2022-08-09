package main

import (
	"context"
	"fmt"
	"github.com/dtc03012/me/file_server/api"
	"github.com/dtc03012/me/file_server/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"time"

	"google.golang.org/api/drive/v3"
)

const (
	MaxUploadSize  = 10 * 1024 * 1024 // 1MB
	filePortNumber = "8282"
)

var (
	ctx      context.Context
	driveSrv *drive.Service
	fileSrv  *handler.FileServer
)

type Response struct {
	filename string
}

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	ext := filepath.Ext(fileHeader.Filename)
	if ext != ".png" && ext != ".jpg" && ext != ".jpeg" {
		http.Error(w, "Not supported extension of the file", http.StatusBadRequest)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, MaxUploadSize)
	if err := r.ParseMultipartForm(MaxUploadSize); err != nil {
		//http.Error(w, "The uploaded file is too big. Please choose an file that's less than 1MB in size", http.StatusBadRequest)
		return
	}

	uploadFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename))

	//// Create the uploads folder if it doesn't
	//// already exist
	//err = os.MkdirAll("./uploads", os.ModePerm)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//// Create a new file in the uploads directory
	//
	//uploadFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename))
	//dst, err := os.Create(fmt.Sprintf("./uploads/%s", uploadFileName))
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//defer dst.Close()
	//
	//// Copy the uploaded file to the filesystem
	//// at the specified destination
	//_, err = io.Copy(dst, file)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	err = fileSrv.UploadFileToDrive(ctx, driveSrv, file, uploadFileName)
	if err != nil {
		return
	}

	w.Write([]byte(fmt.Sprintf("{\"filename\": \"%s\"}", uploadFileName)))
}

func GetFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	params := mux.Vars(r)
	fileName := params["fileName"]

	file, err := fileSrv.GetFileFromDrive(ctx, driveSrv, fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(file)
}

func startFileServer(ctx context.Context) {

	var err error

	driveSrv, err = api.GoogleDriveAPIInit(ctx)
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
		return
	}

	fileSrv = handler.NewFileServer()

	mux := mux.NewRouter()
	mux.HandleFunc("/file/upload-file", UploadFileHandler)
	mux.HandleFunc("/file/get-file/{fileName}", GetFileHandler)

	log.Printf("start File server on %s port", filePortNumber)
	if err := http.ListenAndServe(":4500", mux); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func main() {

	ctx = context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		startFileServer(ctx)
		wg.Done()
	}()

	wg.Wait()
}
