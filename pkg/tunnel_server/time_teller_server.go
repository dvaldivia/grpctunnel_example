/*
Copyright © 2023 Daniel Valdivia <c@dvm.sh>
*/
package tunnel_server

import (
	"context"
	pb "github.com/dvaldivia/grpctunnel_example/stubs"
	"log"
	"time"
)

type MyTimeTellServer struct {
	pb.UnimplementedTimeTellerServer
}

func (m *MyTimeTellServer) WhatTimeIsIt(ctx context.Context, request *pb.TellRequest) (*pb.TimeResponse, error) {
	log.Println("WhatTimeIsIt called")
	return &pb.TimeResponse{
		Message: time.Now().String(),
	}, nil
}
