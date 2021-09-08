package service

import (
	tasks3pb "github.com/Saser/pdp/tasks3/tasks3_go_proto"
)

type Server struct {
	tasks3pb.UnimplementedTasksServer

	tasks []*tasks3pb.Task
}

func NewServer() *Server {
	return &Server{}
}
