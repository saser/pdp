package server

import (
	"context"
	"testing"

	"github.com/Saser/pdp/testing/grpctest"

	taskspb "github.com/Saser/pdp/tasks2/tasks_go_proto"
)

func taskLessFunc(t1, t2 *taskspb.Task) bool {
	return t1.GetName() < t2.GetName()
}

func labelLessFunc(l1, l2 *taskspb.Label) bool {
	return l1.GetName() < l2.GetName()
}

type testTasksClient struct {
	taskspb.TasksClient
}

func (c testTasksClient) GetTaskT(ctx context.Context, t *testing.T, req *taskspb.GetTaskRequest) *taskspb.Task {
	t.Helper()
	task, err := c.GetTask(ctx, req)
	if err != nil {
		t.Fatalf("GetTask(%v) err = %v; want nil", req, err)
	}
	return task
}

func (c testTasksClient) ListTasksT(ctx context.Context, t *testing.T, req *taskspb.ListTasksRequest) *taskspb.ListTasksResponse {
	t.Helper()
	res, err := c.ListTasks(ctx, req)
	if err != nil {
		t.Fatalf("ListTasks(%v) err = %v; want nil", req, err)
	}
	return res
}

func (c testTasksClient) CreateTaskT(ctx context.Context, t *testing.T, req *taskspb.CreateTaskRequest) *taskspb.Task {
	t.Helper()
	task, err := c.CreateTask(ctx, req)
	if err != nil {
		t.Fatalf("CreateTask(%v) err = %v; want nil", req, err)
	}
	return task
}

func (c testTasksClient) UpdateTaskT(ctx context.Context, t *testing.T, req *taskspb.UpdateTaskRequest) *taskspb.Task {
	t.Helper()
	task, err := c.UpdateTask(ctx, req)
	if err != nil {
		t.Fatalf("UpdateTask(%v) err = %v; want nil", req, err)
	}
	return task
}

func (c testTasksClient) AddDependencyT(ctx context.Context, t *testing.T, req *taskspb.AddDependencyRequest) *taskspb.Task {
	t.Helper()
	task, err := c.AddDependency(ctx, req)
	if err != nil {
		t.Fatalf("AddDependency(%v) err = %v; want nil", req, err)
	}
	return task
}

func (c testTasksClient) RemoveDependencyT(ctx context.Context, t *testing.T, req *taskspb.RemoveDependencyRequest) *taskspb.Task {
	t.Helper()
	task, err := c.RemoveDependency(ctx, req)
	if err != nil {
		t.Fatalf("RemoveDependency(%v) err = %v; want nil", req, err)
	}
	return task
}

func (c testTasksClient) AddLabelT(ctx context.Context, t *testing.T, req *taskspb.AddLabelRequest) *taskspb.Task {
	t.Helper()
	task, err := c.AddLabel(ctx, req)
	if err != nil {
		t.Fatalf("AddLabel(%v) err = %v; want nil", req, err)
	}
	return task
}

func (c testTasksClient) RemoveLabelT(ctx context.Context, t *testing.T, req *taskspb.RemoveLabelRequest) *taskspb.Task {
	t.Helper()
	task, err := c.RemoveLabel(ctx, req)
	if err != nil {
		t.Fatalf("RemoveLabel(%v) err = %v; want nil", req, err)
	}
	return task
}

func (c testTasksClient) ListLabelsT(ctx context.Context, t *testing.T, req *taskspb.ListLabelsRequest) *taskspb.ListLabelsResponse {
	t.Helper()
	res, err := c.ListLabels(ctx, req)
	if err != nil {
		t.Fatalf("ListLabels(%v) err = %v; want nil", req, err)
	}
	return res
}

func (c testTasksClient) CreateLabelT(ctx context.Context, t *testing.T, req *taskspb.CreateLabelRequest) *taskspb.Label {
	t.Helper()
	label, err := c.CreateLabel(ctx, req)
	if err != nil {
		t.Fatalf("CreateLabel(%v) err = %v; want nil", req, err)
	}
	return label
}

func setup(t *testing.T) testTasksClient {
	t.Helper()
	cc := grpctest.NewClientConnT(t, &taskspb.Tasks_ServiceDesc, New())
	return testTasksClient{TasksClient: taskspb.NewTasksClient(cc)}
}
