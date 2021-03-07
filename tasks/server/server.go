package server

import (
	taskspb "github.com/Saser/pdp/tasks/tasks_go_proto"
)

type Server struct {
	taskspb.UnimplementedTasksServer
}

func New() *Server {
	return &Server{}
}
