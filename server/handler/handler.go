package handler

import (
	"github.com/dtc03012/me/db/service"
	pb "github.com/dtc03012/me/protobuf/proto/service"
)

type MeServer struct {
	pb.MeServer
	db service.DBService
}

func NewMeServer() *MeServer {
	meServer := &MeServer{
		db: service.NewDBService(),
	}
	return meServer
}
