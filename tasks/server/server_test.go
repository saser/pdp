package server

import (
	"context"
	"testing"

	"github.com/Saser/pdp/testing/errtest"
	"github.com/Saser/pdp/testing/grpctest"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/testing/protocmp"

	taskspb "github.com/Saser/pdp/tasks/tasks_go_proto"
)

func less(t1, t2 *taskspb.Task) bool {
	return t1.GetName() < t2.GetName()
}

type testTasksClient struct {
	taskspb.TasksClient
}

func setup(t *testing.T) testTasksClient {
	t.Helper()
	register := func(s *grpc.Server) { taskspb.RegisterTasksServer(s, New()) }
	cc := grpctest.NewClientConnT(t, register)
	return testTasksClient{
		TasksClient: taskspb.NewTasksClient(cc),
	}
}

func (c testTasksClient) CreateTaskT(ctx context.Context, t *testing.T, req *taskspb.CreateTaskRequest) *taskspb.Task {
	t.Helper()
	task, err := c.CreateTask(ctx, req)
	if err != nil {
		t.Errorf("CreateTask(%v) err = %v; want nil", req, err)
	}
	return task
}

func (c testTasksClient) DeleteTaskT(ctx context.Context, t *testing.T, req *taskspb.DeleteTaskRequest) *taskspb.Task {
	t.Helper()
	task, err := c.DeleteTask(ctx, req)
	if err != nil {
		t.Errorf("DeleteTask(%v) err = %v; want nil", req, err)
	}
	return task
}

func Test_setup(t *testing.T) {
	_ = setup(t)
}

func TestGetTask_OK(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	want := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "My task",
		},
	})
	req := &taskspb.GetTaskRequest{
		Name: want.GetName(),
	}
	got, err := c.GetTask(ctx, req)
	if err != nil {
		t.Errorf("GetTask(%v) err = %v; want nil", req, err)
	}
	if diff := cmp.Diff(want, got, protocmp.Transform()); diff != "" {
		t.Errorf("difference between CreateTask() and GetTask() (-want +got)\n%s", diff)
	}
}

func TestGetTask_Errors(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "Some task",
		},
	})
	for _, tt := range []struct {
		name string
		req  *taskspb.GetTaskRequest
		tf   errtest.TestFunc
	}{
		{
			name: "EmptyName",
			req: &taskspb.GetTaskRequest{
				Name: "",
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("empty name"),
			),
		},
		{
			name: "InvalidName",
			req: &taskspb.GetTaskRequest{
				Name: "issues/1",
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("issues/1"),
				errtest.ErrorContains("invalid task name"),
				errtest.ErrorContains("tasks/{task}"),
			),
		},
		{
			name: "NonExistentTask",
			req: &taskspb.GetTaskRequest{
				Name: "tasks/2",
			},
			tf: errtest.All(
				grpctest.WantCode(codes.NotFound),
				errtest.ErrorContains("tasks/2"),
			),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.GetTask(ctx, tt.req)
			tt.tf(t, err)
		})
	}
}

func TestGetTask_AfterDeletion(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	task := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "Some task",
		},
	})
	deleted := c.DeleteTaskT(ctx, t, &taskspb.DeleteTaskRequest{
		Name: task.GetName(),
	})
	req := &taskspb.GetTaskRequest{
		Name: task.GetName(),
	}
	got, err := c.GetTask(ctx, req)
	if err != nil {
		t.Errorf("GetTask(%v) err = %v; want nil", req, err)
	}
	if diff := cmp.Diff(deleted, got, protocmp.Transform()); diff != "" {
		t.Errorf("diff between DeleteTask() and GetTask() (-want +got)\n%s", diff)
	}
}

func TestListTasks_OK(t *testing.T) {
	ctx := context.Background()
	for _, tt := range []struct {
		name  string
		tasks []*taskspb.Task
	}{
		{
			name:  "NoTasks",
			tasks: nil,
		},
		{
			name: "SomeTasks",
			tasks: []*taskspb.Task{
				{Title: "First task"},
				{Title: "Second task"},
				{Title: "Third task"},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			c := setup(t)
			want := make([]*taskspb.Task, len(tt.tasks))
			for i, task := range tt.tasks {
				want[i] = c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
					Task: task,
				})
			}
			req := &taskspb.ListTasksRequest{}
			res, err := c.ListTasks(ctx, req)
			if err != nil {
				t.Errorf("ListTasks(%v) err = %v; want nil", req, err)
			}
			if res.GetNextPageToken() != "" {
				t.Errorf("res.GetNextPageToken() = %q; want empty", res.GetNextPageToken())
			}
			got := res.GetTasks()
			if diff := cmp.Diff(want, got, protocmp.Transform(), cmpopts.EquateEmpty(), cmpopts.SortSlices(less)); diff != "" {
				t.Errorf("diff between created tasks and listed tasks (-want +got)\n%s", diff)
			}
		})
	}
}

func TestListTasks_Errors(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	for _, tt := range []struct {
		name string
		req  *taskspb.ListTasksRequest
		tf   errtest.TestFunc
	}{
		{
			name: "NegativePageSize",
			req:  &taskspb.ListTasksRequest{PageSize: -1},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("negative page size"),
			),
		},
		{
			name: "InvalidPageToken",
			req:  &taskspb.ListTasksRequest{PageToken: "This is rubbish"},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("invalid page token"),
				errtest.ErrorContains("This is rubbish"),
			),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.ListTasks(ctx, tt.req)
			tt.tf(t, err)
		})
	}
}

func TestListTasks_DeletedTasks(t *testing.T) {
	ctx := context.Background()
	c := setup(t)

	existing := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "Existing task",
		},
	})
	deleted := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "Deleted task",
		},
	})
	deleted = c.DeleteTaskT(ctx, t, &taskspb.DeleteTaskRequest{
		Name: deleted.GetName(),
	})

	for _, tt := range []struct {
		name string
		req  *taskspb.ListTasksRequest
		want []*taskspb.Task
	}{
		{
			name: "NoShowDeleted",
			req:  &taskspb.ListTasksRequest{ShowDeleted: false},
			want: []*taskspb.Task{
				existing,
			},
		},
		{
			name: "ShowDeleted",
			req:  &taskspb.ListTasksRequest{ShowDeleted: true},
			want: []*taskspb.Task{
				existing,
				deleted,
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			res, err := c.ListTasks(ctx, tt.req)
			if err != nil {
				t.Errorf("ListTasks(%v) err = %v; want nil", tt.req, err)
			}
			got := res.GetTasks()
			if diff := cmp.Diff(tt.want, got, cmpopts.SortSlices(less), protocmp.Transform()); diff != "" {
				t.Errorf("diff in ListTasks() (-want +got)\n%s", diff)
			}
		})
	}
}

func TestListTasks_Pagination_NoShowDeleted_NoDeletedTasks(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	want := []*taskspb.Task{
		{Title: "First task"},
		{Title: "Second task"},
	}
	for i, task := range want {
		want[i] = c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
			Task: task,
		})
	}

	var got []*taskspb.Task

	// Get first page of results that should contain exactly one task and a non-empty NextPageToken.
	req1 := &taskspb.ListTasksRequest{
		PageSize: 1,
	}
	res1, err := c.ListTasks(ctx, req1)
	if err != nil {
		t.Errorf("ListTasks(%v) err = %v; want nil", req1, err)
	}
	if tasks := res1.GetTasks(); len(tasks) != 1 {
		t.Errorf("first page of tasks: got %v (len = %v); want len = 1", tasks, len(tasks))
	}
	if tok := res1.GetNextPageToken(); tok == "" {
		t.Error("first page of tasks: next_page_token is empty")
	}
	got = append(got, res1.GetTasks()...)

	// Get second page of results, which should contain the rest of the tasks.
	req2 := &taskspb.ListTasksRequest{
		PageToken: res1.GetNextPageToken(),
	}
	res2, err := c.ListTasks(ctx, req2)
	if err != nil {
		t.Errorf("ListTasks(%v) err = %v; want nil", req2, err)
	}
	if tok := res2.GetNextPageToken(); tok != "" {
		t.Errorf("second page of tasks: next_page_token = %q; want empty string", tok)
	}
	got = append(got, res2.GetTasks()...)

	if diff := cmp.Diff(want, got, protocmp.Transform(), cmpopts.SortSlices(less)); diff != "" {
		t.Errorf("diff between created and listed tasks (-want +got)\n%s", diff)
	}
}

func TestListTasks_Pagination_NoShowDeleted_SomeDeletedTasks(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	existing := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "Existing task",
		},
	})
	deleted := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "Deleted task",
		},
	})
	deleted = c.DeleteTaskT(ctx, t, &taskspb.DeleteTaskRequest{
		Name: deleted.GetName(),
	})
	want := []*taskspb.Task{
		existing,
	}

	var got []*taskspb.Task

	// Get first page of results that should contain exactly one task and a non-empty NextPageToken.
	req := &taskspb.ListTasksRequest{
		PageSize: 1,
	}
	res, err := c.ListTasks(ctx, req)
	if err != nil {
		t.Errorf("ListTasks(%v) err = %v; want nil", req, err)
	}
	if tasks := res.GetTasks(); len(tasks) != 1 {
		t.Errorf("first page of tasks: got %v (len = %v); want len = 1", tasks, len(tasks))
	}
	if tok := res.GetNextPageToken(); tok != "" {
		t.Errorf("first page of tasks: next_page_token = %q; want empty string", tok)
	}
	got = append(got, res.GetTasks()...)

	if diff := cmp.Diff(want, got, protocmp.Transform(), cmpopts.SortSlices(less)); diff != "" {
		t.Errorf("diff between created and listed tasks (-want +got)\n%s", diff)
	}
}

func TestListTasks_Pagination_ShowDeleted(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	existing := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "Existing task",
		},
	})
	deleted := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "Deleted task",
		},
	})
	deleted = c.DeleteTaskT(ctx, t, &taskspb.DeleteTaskRequest{
		Name: deleted.GetName(),
	})
	want := []*taskspb.Task{
		existing,
		deleted,
	}

	var got []*taskspb.Task

	// Get first page of results that should contain exactly one task and a non-empty NextPageToken.
	req1 := &taskspb.ListTasksRequest{
		PageSize:    1,
		ShowDeleted: true,
	}
	res1, err := c.ListTasks(ctx, req1)
	if err != nil {
		t.Errorf("ListTasks(%v) err = %v; want nil", req1, err)
	}
	if tasks := res1.GetTasks(); len(tasks) != 1 {
		t.Errorf("first page of tasks: got %v (len = %v); want len = 1", tasks, len(tasks))
	}
	if tok := res1.GetNextPageToken(); tok == "" {
		t.Error("first page of tasks: next_page_token is empty")
	}
	got = append(got, res1.GetTasks()...)

	// Get second page of results, which should contain the rest of the tasks.
	req2 := &taskspb.ListTasksRequest{
		PageToken:   res1.GetNextPageToken(),
		ShowDeleted: true,
	}
	res2, err := c.ListTasks(ctx, req2)
	if err != nil {
		t.Errorf("ListTasks(%v) err = %v; want nil", req2, err)
	}
	if tok := res2.GetNextPageToken(); tok != "" {
		t.Errorf("second page of tasks: next_page_token = %q; want empty string", tok)
	}
	got = append(got, res2.GetTasks()...)

	if diff := cmp.Diff(want, got, protocmp.Transform(), cmpopts.SortSlices(less)); diff != "" {
		t.Errorf("diff between created and listed tasks (-want +got)\n%s", diff)
	}
}

func TestListTasks_Pagination_DifferentQueryParameters(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	for _, task := range []*taskspb.Task{
		{Title: "First task"},
		{Title: "Second task"},
	} {
		c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
			Task: task,
		})
	}

	req1 := &taskspb.ListTasksRequest{
		PageSize:    1,
		ShowDeleted: false,
	}
	res1, err := c.ListTasks(ctx, req1)
	if err != nil {
		t.Errorf("ListTasks(%v) err = %v; want nil", req1, err)
	}

	req2 := &taskspb.ListTasksRequest{
		PageSize:    5,    // was 1 in previous request, but is allowed to change
		ShowDeleted: true, // was false in previous request, and is not allowed to change
		PageToken:   res1.GetNextPageToken(),
	}
	_, err = c.ListTasks(ctx, req2)
	if err == nil {
		t.Fatalf("ListTasks(%v) err = nil; want non-nil", req2)
	}
	if got, want := status.Code(err), codes.InvalidArgument; got != want {
		t.Errorf("status.Code(%v) = %v; want %v", err, got, want)
	}
}

func TestListTasks_GetTasks_NoDeletedTasks(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	for _, task := range []*taskspb.Task{
		{Title: "First task"},
		{Title: "Second task"},
		{Title: "Third task"},
	} {
		c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
			Task: task,
		})
	}

	listReq := &taskspb.ListTasksRequest{}
	res, err := c.ListTasks(ctx, listReq)
	if err != nil {
		t.Errorf("ListTasks(%v) err = %v; want nil", listReq, err)
	}
	for _, want := range res.GetTasks() {
		getReq := &taskspb.GetTaskRequest{
			Name: want.GetName(),
		}
		got, err := c.GetTask(ctx, getReq)
		if err != nil {
			t.Errorf("GetTask(%v) err = %v; want nil", getReq, err)
		}
		if diff := cmp.Diff(want, got, protocmp.Transform()); diff != "" {
			t.Errorf("difference between ListTasks() and GetTask() (-want +got)\n%s", diff)
		}
	}
}

func TestListTasks_GetTasks_SomeDeletedTasks(t *testing.T) {
	ctx := context.Background()
	c := setup(t)

	c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "Existing task",
		},
	})
	deleted := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "Deleted task",
		},
	})
	c.DeleteTaskT(ctx, t, &taskspb.DeleteTaskRequest{
		Name: deleted.GetName(),
	})

	listReq := &taskspb.ListTasksRequest{
		ShowDeleted: true,
	}
	res, err := c.ListTasks(ctx, listReq)
	if err != nil {
		t.Errorf("ListTasks(%v) err = %v; want nil", listReq, err)
	}
	for _, want := range res.GetTasks() {
		getReq := &taskspb.GetTaskRequest{
			Name: want.GetName(),
		}
		got, err := c.GetTask(ctx, getReq)
		if err != nil {
			t.Errorf("GetTask(%v) err = %v; want nil", getReq, err)
		}
		if diff := cmp.Diff(want, got, protocmp.Transform()); diff != "" {
			t.Errorf("difference between ListTasks() and GetTask() (-want +got)\n%s", diff)
		}
	}
}
