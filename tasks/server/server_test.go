package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/Saser/pdp/testing/errtest"
	"github.com/Saser/pdp/testing/grpctest"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

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

func (c testTasksClient) GetTaskT(ctx context.Context, t *testing.T, req *taskspb.GetTaskRequest) *taskspb.Task {
	t.Helper()
	task, err := c.GetTask(ctx, req)
	if err != nil {
		t.Errorf("GetTask(%v) err = %v; want nil", req, err)
	}
	return task
}

func (c testTasksClient) ListTasksT(ctx context.Context, t *testing.T, req *taskspb.ListTasksRequest) *taskspb.ListTasksResponse {
	t.Helper()
	res, err := c.ListTasks(ctx, req)
	if err != nil {
		t.Errorf("ListTasks(%v) err = %v; want nil", req, err)
	}
	return res
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

func TestCreateTask_OK(t *testing.T) {
	ctx := context.Background()
	c := setup(t)

	req := &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "My task",
		},
	}
	if _, err := c.CreateTask(ctx, req); err != nil {
		t.Errorf("CreateTask(%v) err = %v; want nil", req, err)
	}
}

func TestCreateTask_Errors(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	for _, tt := range []struct {
		name string
		req  *taskspb.CreateTaskRequest
		tf   errtest.TestFunc
	}{
		{
			name: "EmptyTitle",
			req: &taskspb.CreateTaskRequest{
				Task: &taskspb.Task{
					Title: "",
				},
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("empty title"),
			),
		},
		{
			name: "NonEmptyName",
			req: &taskspb.CreateTaskRequest{
				Task: &taskspb.Task{
					Name:  "tasks/1",
					Title: "My task",
				},
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains(`"name" must be empty`),
				errtest.ErrorContains(`"tasks/1"`),
			),
		},
		{
			name: "DeletedTrue",
			req: &taskspb.CreateTaskRequest{
				Task: &taskspb.Task{
					Deleted: true,
					Title:   "My task",
				},
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains(`"deleted" must be false`),
			),
		},
		{
			name: "CompletedTrue",
			req: &taskspb.CreateTaskRequest{
				Task: &taskspb.Task{
					Title:     "My task",
					Completed: true,
				},
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains(`"completed" must be false`),
				errtest.ErrorContains("SetCompleted"),
			),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.CreateTask(ctx, tt.req)
			if err == nil {
				t.Fatalf("CreateTask(%v) err = nil; want non-nil", tt.req)
			}
			tt.tf(t, err)
			if t.Failed() {
				t.Logf("request: %v", tt.req)
			}
		})
	}
}

func TestUpdateTask_OK(t *testing.T) {
	ctx := context.Background()
	for _, tt := range []struct {
		name      string
		createReq *taskspb.CreateTaskRequest
		updateReq *taskspb.UpdateTaskRequest
		want      *taskspb.Task
	}{
		{
			name: "UpdateTitle_NilMask",
			createReq: &taskspb.CreateTaskRequest{
				Task: &taskspb.Task{
					Title: "Title before",
				},
			},
			updateReq: &taskspb.UpdateTaskRequest{
				Task: &taskspb.Task{
					Title: "Title after",
				},
			},
			want: &taskspb.Task{
				Title: "Title after",
			},
		},
		{
			name: "UpdateTitle_StarMask",
			createReq: &taskspb.CreateTaskRequest{
				Task: &taskspb.Task{
					Title: "Title before",
				},
			},
			updateReq: &taskspb.UpdateTaskRequest{
				Task: &taskspb.Task{
					Title: "Title after",
				},
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{"*"},
				},
			},
			want: &taskspb.Task{
				Title: "Title after",
			},
		},
		{
			name: "UpdateTitle_NonEmptyMask",
			createReq: &taskspb.CreateTaskRequest{
				Task: &taskspb.Task{
					Title: "Title before",
				},
			},
			updateReq: &taskspb.UpdateTaskRequest{
				Task: &taskspb.Task{
					Title: "Title after",
				},
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{"title"},
				},
			},
			want: &taskspb.Task{
				Title: "Title after",
			},
		},
		{
			name: "IgnoreOutputOnly_StarMask",
			createReq: &taskspb.CreateTaskRequest{
				Task: &taskspb.Task{
					Title: "Title before",
				},
			},
			updateReq: &taskspb.UpdateTaskRequest{
				Task: &taskspb.Task{
					Title:     "Title after",
					Deleted:   true,
					Completed: true,
				},
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{"*"},
				},
			},
			want: &taskspb.Task{
				Title:     "Title after",
				Deleted:   false, // specified just for clarity
				Completed: false, // specified just for clarity
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			c := setup(t)
			task := c.CreateTaskT(ctx, t, tt.createReq)
			tt.updateReq.GetTask().Name = task.GetName()
			got, err := c.UpdateTask(ctx, tt.updateReq)
			if err != nil {
				t.Fatalf("UpdateTask(%v) err = %v; want nil", tt.updateReq, err)
			}
			if diff := cmp.Diff(tt.want, got, protocmp.Transform(), protocmp.IgnoreFields(&taskspb.Task{}, "name")); diff != "" {
				t.Errorf("unexpected result of Update (-want +got)\n%s", diff)
			}
		})
	}
}

func TestUpdateTask_Errors(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	task := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "Some task",
		},
	})
	clone := func(modify func(task *taskspb.Task)) *taskspb.Task {
		task2 := proto.Clone(task).(*taskspb.Task)
		modify(task2)
		return task2
	}
	for _, tt := range []struct {
		name string
		req  *taskspb.UpdateTaskRequest
		tf   errtest.TestFunc
	}{
		{
			name: "EmptyName",
			req: &taskspb.UpdateTaskRequest{
				Task: &taskspb.Task{
					Name: "", // explicitly set to zero value for clarity
				},
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("empty name"),
			),
		},
		{
			name: "InvalidName",
			req: &taskspb.UpdateTaskRequest{
				Task: &taskspb.Task{
					Name: "invalidname/1",
				},
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("invalid name"),
			),
		},
		{
			name: "NonExistentTask",
			req: &taskspb.UpdateTaskRequest{
				Task: &taskspb.Task{
					Name: "tasks/999",
				},
			},
			tf: errtest.All(
				grpctest.WantCode(codes.NotFound),
				errtest.ErrorContains(`"tasks/999"`),
			),
		},
		{
			name: "EmptyTitle_DirectFieldMask",
			req: &taskspb.UpdateTaskRequest{
				Task:       clone(func(task *taskspb.Task) { task.Title = "" }),
				UpdateMask: &fieldmaskpb.FieldMask{Paths: []string{"title"}},
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("empty title"),
			),
		},
		{
			name: "EmptyTitle_StarMask",
			req: &taskspb.UpdateTaskRequest{
				Task:       clone(func(task *taskspb.Task) { task.Title = "" }),
				UpdateMask: &fieldmaskpb.FieldMask{Paths: []string{"*"}},
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("empty title"),
			),
		},
		{
			name: "OutputOnlyFieldsInMask",
			req: &taskspb.UpdateTaskRequest{
				Task: clone(func(task *taskspb.Task) {
					task.Deleted = true
					task.Completed = true
				}),
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"deleted",
						"completed",
					},
				},
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("output only"),
				errtest.ErrorContains(`"deleted"`),
				errtest.ErrorContains(`"completed"`),
			),
		},
		{
			name: "InvalidPathsInFieldMask",
			req: &taskspb.UpdateTaskRequest{
				Task: clone(func(task *taskspb.Task) { task.Title = "Some other title" }),
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"foobar",
						"baz.quux",
					},
				},
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("invalid path"),
				errtest.ErrorContains(`"foobar"`),
				errtest.ErrorContains(`"baz.quux"`),
			),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.UpdateTask(ctx, tt.req)
			tt.tf(t, err)

			// Since we expect the update to fail, we verify that the task was not
			// actually modified.
			after, err := c.GetTask(ctx, &taskspb.GetTaskRequest{
				Name: task.GetName(),
			})
			if diff := cmp.Diff(task, after, protocmp.Transform()); diff != "" {
				t.Errorf("resource updated even with errors (-before +after)\n%s", diff)
			}
		})
	}
}

func TestDeleteTask(t *testing.T) {
	ctx := context.Background()
	c := setup(t)

	// Create a task that should later be deleted.
	task := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "to be deleted",
		},
	})

	// Verify that the task can be deleted.
	req := &taskspb.DeleteTaskRequest{
		Name: task.GetName(),
	}
	deleted, err := c.DeleteTask(ctx, req)
	if err != nil {
		t.Fatalf("DeleteTask(%v) err = %v; want nil", req, err)
	}
	if deleted.GetDeleted() == false {
		t.Error("deleted.GetDeleted() = false; want true")
	}
}

func TestDeleteTask_ThenGetTask(t *testing.T) {
	ctx := context.Background()
	c := setup(t)

	// Create a task then delete it.
	task := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "to be deleted",
		},
	})
	task = c.DeleteTaskT(ctx, t, &taskspb.DeleteTaskRequest{
		Name: task.GetName(),
	})

	// Verify that the task can still be queried using GetTask.
	got := c.GetTaskT(ctx, t, &taskspb.GetTaskRequest{
		Name: task.GetName(),
	})
	if diff := cmp.Diff(task, got, protocmp.Transform()); diff != "" {
		t.Errorf("GetTask() after deleting has diff (-want +got)\n%s", diff)
	}
}

func TestDeleteTask_ThenListTasks(t *testing.T) {
	ctx := context.Background()
	c := setup(t)

	// Create a task then delete it.
	deleted := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "to be deleted",
		},
	})
	deleted = c.DeleteTaskT(ctx, t, &taskspb.DeleteTaskRequest{
		Name: deleted.GetName(),
	})

	// Create another task that should not be deleted.
	notDeleted := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "should not be deleted",
		},
	})

	for _, tt := range []struct {
		name string
		req  *taskspb.ListTasksRequest
		want []*taskspb.Task
	}{
		{
			name: "ShowDeleted",
			req: &taskspb.ListTasksRequest{
				ShowDeleted: true,
			},
			want: []*taskspb.Task{
				deleted,
				notDeleted,
			},
		},
		{
			name: "NoShowDeleted",
			req: &taskspb.ListTasksRequest{
				ShowDeleted: false,
			},
			want: []*taskspb.Task{
				notDeleted,
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			res := c.ListTasksT(ctx, t, tt.req)
			if diff := cmp.Diff(tt.want, res.GetTasks(), protocmp.Transform(), cmpopts.SortSlices(less)); diff != "" {
				t.Errorf("unexpected result from listing tasks (-want +got)\n%s", diff)
			}
		})
	}
}

func TestDeleteTask_ThenDeleteAgain(t *testing.T) {
	ctx := context.Background()
	c := setup(t)

	// Create a task then delete it.
	deleted := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "to be deleted",
		},
	})
	req := &taskspb.DeleteTaskRequest{
		Name: deleted.GetName(),
	}
	deleted = c.DeleteTaskT(ctx, t, req)

	// Repeating the request should fail.
	_, err := c.DeleteTask(ctx, req)
	errtest.All(
		grpctest.WantCode(codes.InvalidArgument),
		errtest.ErrorContains("already deleted"),
		errtest.ErrorContains(fmt.Sprintf("%q", deleted.GetName())),
	)(t, err)
}

func TestDeleteTask_Errors(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	for _, tt := range []struct {
		name string
		req  *taskspb.DeleteTaskRequest
		tf   errtest.TestFunc
	}{
		{
			name: "EmptyName",
			req: &taskspb.DeleteTaskRequest{
				Name: "", // explicitly set for clarity
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("empty name"),
			),
		},
		{
			name: "InvalidName",
			req: &taskspb.DeleteTaskRequest{
				Name: "invalidname",
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("invalid name"),
				errtest.ErrorContains(`"invalidname"`),
			),
		},
		{
			name: "NonExistentTask",
			req: &taskspb.DeleteTaskRequest{
				Name: "tasks/999",
			},
			tf: errtest.All(
				grpctest.WantCode(codes.NotFound),
				errtest.ErrorContains(`"tasks/999"`),
			),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.DeleteTask(ctx, tt.req)
			tt.tf(t, err)
		})
	}
}
