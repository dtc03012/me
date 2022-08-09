package handler

import "github.com/dtc03012/me/file_server/db/service"

type FileServer struct {
	db service.DBService
}

func NewFileServer() *FileServer {
	fileServer := &FileServer{
		db: service.NewDBService(),
	}
	return fileServer
}
