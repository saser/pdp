// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package tasks_go_proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TasksClient is the client API for Tasks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TasksClient interface {
	GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*Task, error)
	ListTasks(ctx context.Context, in *ListTasksRequest, opts ...grpc.CallOption) (*ListTasksResponse, error)
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*Task, error)
	UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*Task, error)
	AddComment(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*Task, error)
	AddDependency(ctx context.Context, in *AddDependencyRequest, opts ...grpc.CallOption) (*Task, error)
	RemoveDependency(ctx context.Context, in *RemoveDependencyRequest, opts ...grpc.CallOption) (*Task, error)
	AddLabel(ctx context.Context, in *AddLabelRequest, opts ...grpc.CallOption) (*Task, error)
	RemoveLabel(ctx context.Context, in *RemoveLabelRequest, opts ...grpc.CallOption) (*Task, error)
	CompleteTask(ctx context.Context, in *CompleteTaskRequest, opts ...grpc.CallOption) (*Task, error)
	UncompleteTask(ctx context.Context, in *UncompleteTaskRequest, opts ...grpc.CallOption) (*Task, error)
	GetEvent(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*Event, error)
	ListEvents(ctx context.Context, in *ListEventsRequest, opts ...grpc.CallOption) (*ListEventsResponse, error)
	GetLabel(ctx context.Context, in *GetLabelRequest, opts ...grpc.CallOption) (*Label, error)
	ListLabels(ctx context.Context, in *ListLabelsRequest, opts ...grpc.CallOption) (*ListLabelsResponse, error)
	CreateLabel(ctx context.Context, in *CreateLabelRequest, opts ...grpc.CallOption) (*Label, error)
	UpdateLabel(ctx context.Context, in *UpdateLabelRequest, opts ...grpc.CallOption) (*Label, error)
}

type tasksClient struct {
	cc grpc.ClientConnInterface
}

func NewTasksClient(cc grpc.ClientConnInterface) TasksClient {
	return &tasksClient{cc}
}

func (c *tasksClient) GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/tasks2.Tasks/GetTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) ListTasks(ctx context.Context, in *ListTasksRequest, opts ...grpc.CallOption) (*ListTasksResponse, error) {
	out := new(ListTasksResponse)
	err := c.cc.Invoke(ctx, "/tasks2.Tasks/ListTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/tasks2.Tasks/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/tasks2.Tasks/UpdateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) AddComment(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/tasks2.Tasks/AddComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) AddDependency(ctx context.Context, in *AddDependencyRequest, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/tasks2.Tasks/AddDependency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) RemoveDependency(ctx context.Context, in *RemoveDependencyRequest, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/tasks2.Tasks/RemoveDependency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) AddLabel(ctx context.Context, in *AddLabelRequest, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/tasks2.Tasks/AddLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) RemoveLabel(ctx context.Context, in *RemoveLabelRequest, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/tasks2.Tasks/RemoveLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) CompleteTask(ctx context.Context, in *CompleteTaskRequest, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/tasks2.Tasks/CompleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) UncompleteTask(ctx context.Context, in *UncompleteTaskRequest, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/tasks2.Tasks/UncompleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) GetEvent(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/tasks2.Tasks/GetEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) ListEvents(ctx context.Context, in *ListEventsRequest, opts ...grpc.CallOption) (*ListEventsResponse, error) {
	out := new(ListEventsResponse)
	err := c.cc.Invoke(ctx, "/tasks2.Tasks/ListEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) GetLabel(ctx context.Context, in *GetLabelRequest, opts ...grpc.CallOption) (*Label, error) {
	out := new(Label)
	err := c.cc.Invoke(ctx, "/tasks2.Tasks/GetLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) ListLabels(ctx context.Context, in *ListLabelsRequest, opts ...grpc.CallOption) (*ListLabelsResponse, error) {
	out := new(ListLabelsResponse)
	err := c.cc.Invoke(ctx, "/tasks2.Tasks/ListLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) CreateLabel(ctx context.Context, in *CreateLabelRequest, opts ...grpc.CallOption) (*Label, error) {
	out := new(Label)
	err := c.cc.Invoke(ctx, "/tasks2.Tasks/CreateLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) UpdateLabel(ctx context.Context, in *UpdateLabelRequest, opts ...grpc.CallOption) (*Label, error) {
	out := new(Label)
	err := c.cc.Invoke(ctx, "/tasks2.Tasks/UpdateLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TasksServer is the server API for Tasks service.
// All implementations must embed UnimplementedTasksServer
// for forward compatibility
type TasksServer interface {
	GetTask(context.Context, *GetTaskRequest) (*Task, error)
	ListTasks(context.Context, *ListTasksRequest) (*ListTasksResponse, error)
	CreateTask(context.Context, *CreateTaskRequest) (*Task, error)
	UpdateTask(context.Context, *UpdateTaskRequest) (*Task, error)
	AddComment(context.Context, *AddCommentRequest) (*Task, error)
	AddDependency(context.Context, *AddDependencyRequest) (*Task, error)
	RemoveDependency(context.Context, *RemoveDependencyRequest) (*Task, error)
	AddLabel(context.Context, *AddLabelRequest) (*Task, error)
	RemoveLabel(context.Context, *RemoveLabelRequest) (*Task, error)
	CompleteTask(context.Context, *CompleteTaskRequest) (*Task, error)
	UncompleteTask(context.Context, *UncompleteTaskRequest) (*Task, error)
	GetEvent(context.Context, *GetEventRequest) (*Event, error)
	ListEvents(context.Context, *ListEventsRequest) (*ListEventsResponse, error)
	GetLabel(context.Context, *GetLabelRequest) (*Label, error)
	ListLabels(context.Context, *ListLabelsRequest) (*ListLabelsResponse, error)
	CreateLabel(context.Context, *CreateLabelRequest) (*Label, error)
	UpdateLabel(context.Context, *UpdateLabelRequest) (*Label, error)
	mustEmbedUnimplementedTasksServer()
}

// UnimplementedTasksServer must be embedded to have forward compatible implementations.
type UnimplementedTasksServer struct {
}

func (UnimplementedTasksServer) GetTask(context.Context, *GetTaskRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (UnimplementedTasksServer) ListTasks(context.Context, *ListTasksRequest) (*ListTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTasks not implemented")
}
func (UnimplementedTasksServer) CreateTask(context.Context, *CreateTaskRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedTasksServer) UpdateTask(context.Context, *UpdateTaskRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}
func (UnimplementedTasksServer) AddComment(context.Context, *AddCommentRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (UnimplementedTasksServer) AddDependency(context.Context, *AddDependencyRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDependency not implemented")
}
func (UnimplementedTasksServer) RemoveDependency(context.Context, *RemoveDependencyRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDependency not implemented")
}
func (UnimplementedTasksServer) AddLabel(context.Context, *AddLabelRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLabel not implemented")
}
func (UnimplementedTasksServer) RemoveLabel(context.Context, *RemoveLabelRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveLabel not implemented")
}
func (UnimplementedTasksServer) CompleteTask(context.Context, *CompleteTaskRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteTask not implemented")
}
func (UnimplementedTasksServer) UncompleteTask(context.Context, *UncompleteTaskRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UncompleteTask not implemented")
}
func (UnimplementedTasksServer) GetEvent(context.Context, *GetEventRequest) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (UnimplementedTasksServer) ListEvents(context.Context, *ListEventsRequest) (*ListEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEvents not implemented")
}
func (UnimplementedTasksServer) GetLabel(context.Context, *GetLabelRequest) (*Label, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLabel not implemented")
}
func (UnimplementedTasksServer) ListLabels(context.Context, *ListLabelsRequest) (*ListLabelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLabels not implemented")
}
func (UnimplementedTasksServer) CreateLabel(context.Context, *CreateLabelRequest) (*Label, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLabel not implemented")
}
func (UnimplementedTasksServer) UpdateLabel(context.Context, *UpdateLabelRequest) (*Label, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLabel not implemented")
}
func (UnimplementedTasksServer) mustEmbedUnimplementedTasksServer() {}

// UnsafeTasksServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TasksServer will
// result in compilation errors.
type UnsafeTasksServer interface {
	mustEmbedUnimplementedTasksServer()
}

func RegisterTasksServer(s grpc.ServiceRegistrar, srv TasksServer) {
	s.RegisterService(&Tasks_ServiceDesc, srv)
}

func _Tasks_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasks2.Tasks/GetTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).GetTask(ctx, req.(*GetTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_ListTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).ListTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasks2.Tasks/ListTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).ListTasks(ctx, req.(*ListTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasks2.Tasks/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasks2.Tasks/UpdateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).UpdateTask(ctx, req.(*UpdateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasks2.Tasks/AddComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).AddComment(ctx, req.(*AddCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_AddDependency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDependencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).AddDependency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasks2.Tasks/AddDependency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).AddDependency(ctx, req.(*AddDependencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_RemoveDependency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveDependencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).RemoveDependency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasks2.Tasks/RemoveDependency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).RemoveDependency(ctx, req.(*RemoveDependencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_AddLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).AddLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasks2.Tasks/AddLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).AddLabel(ctx, req.(*AddLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_RemoveLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).RemoveLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasks2.Tasks/RemoveLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).RemoveLabel(ctx, req.(*RemoveLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_CompleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).CompleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasks2.Tasks/CompleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).CompleteTask(ctx, req.(*CompleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_UncompleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UncompleteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).UncompleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasks2.Tasks/UncompleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).UncompleteTask(ctx, req.(*UncompleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasks2.Tasks/GetEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).GetEvent(ctx, req.(*GetEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_ListEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).ListEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasks2.Tasks/ListEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).ListEvents(ctx, req.(*ListEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_GetLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).GetLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasks2.Tasks/GetLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).GetLabel(ctx, req.(*GetLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_ListLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).ListLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasks2.Tasks/ListLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).ListLabels(ctx, req.(*ListLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_CreateLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).CreateLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasks2.Tasks/CreateLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).CreateLabel(ctx, req.(*CreateLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_UpdateLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).UpdateLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasks2.Tasks/UpdateLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).UpdateLabel(ctx, req.(*UpdateLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Tasks_ServiceDesc is the grpc.ServiceDesc for Tasks service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tasks_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tasks2.Tasks",
	HandlerType: (*TasksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTask",
			Handler:    _Tasks_GetTask_Handler,
		},
		{
			MethodName: "ListTasks",
			Handler:    _Tasks_ListTasks_Handler,
		},
		{
			MethodName: "CreateTask",
			Handler:    _Tasks_CreateTask_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _Tasks_UpdateTask_Handler,
		},
		{
			MethodName: "AddComment",
			Handler:    _Tasks_AddComment_Handler,
		},
		{
			MethodName: "AddDependency",
			Handler:    _Tasks_AddDependency_Handler,
		},
		{
			MethodName: "RemoveDependency",
			Handler:    _Tasks_RemoveDependency_Handler,
		},
		{
			MethodName: "AddLabel",
			Handler:    _Tasks_AddLabel_Handler,
		},
		{
			MethodName: "RemoveLabel",
			Handler:    _Tasks_RemoveLabel_Handler,
		},
		{
			MethodName: "CompleteTask",
			Handler:    _Tasks_CompleteTask_Handler,
		},
		{
			MethodName: "UncompleteTask",
			Handler:    _Tasks_UncompleteTask_Handler,
		},
		{
			MethodName: "GetEvent",
			Handler:    _Tasks_GetEvent_Handler,
		},
		{
			MethodName: "ListEvents",
			Handler:    _Tasks_ListEvents_Handler,
		},
		{
			MethodName: "GetLabel",
			Handler:    _Tasks_GetLabel_Handler,
		},
		{
			MethodName: "ListLabels",
			Handler:    _Tasks_ListLabels_Handler,
		},
		{
			MethodName: "CreateLabel",
			Handler:    _Tasks_CreateLabel_Handler,
		},
		{
			MethodName: "UpdateLabel",
			Handler:    _Tasks_UpdateLabel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tasks2/tasks.proto",
}
