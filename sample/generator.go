package sample

import (
	pb "go-postgres/proto/user"
)

func NewUser() *pb.DataResponse {
	DataResponse := &pb.DataResponse{
		User: randomUser(),
		Id:   randomId(),
	}
	return DataResponse
}
