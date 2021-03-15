# Test cases for the Tasks API

## `GetTask`

* [x] Empty `name`: fail with `InvalidArgument`
* [x] Invalid `name` (i.e., wrong format): fail with `InvalidArgument`
* [x] Non-existent task: fail with `NotFound`
* [x] Deleted task: succeed!

## `ListTasks`

* [x] No created tasks
* [x] Some created tasks
* [x] Deleted tasks (`show_deleted = false`)
* [x] Deleted tasks (`show_deleted = true`)
* [ ] Each task can also be queried with `GetTask`
* [x] Negative page size: fail with `InvalidArgument`.
* [x] Correct pagination with `show_deleted = false`
    * [x] 2 tasks, set `page_size = 1`, and do two calls
	* [x] 2 tasks, 1 deleted, `page_size = 1` should succeed and return empty `next_page_token`
* [x] Correct pagination with `show_deleted = true`
	* [x] 2 tasks, 1 deleted, `page_size = 1` should succeed and return non-empty `next_page_token`.
* [ ] Some rubbish `page_token`: should fail with `InvalidArgument`.
* [ ] A correct page token but other parameters changed (apart from `page_size`): should fail with `InvalidArgument`.
    * [ ] The only other parameter is `show_deleted` at the moment, but for example `filter` can be added in the future.

## `CreateTask`

* [ ] With missing title (should fail with `InvalidArgument`)
* [ ] With a `name` set (should fail with `InvalidArgument`)
    * [ ] Why? Because a client could think the `name` can be set when creating, and if the name is silently ignored and replaced the client could continue using their own name (if they were to ignore the returned task, which they shouldn't, but yeah).
* [ ] With `deleted = true` (should fail with `InvalidArgument`)
* [ ] With `completed = true` (should fail with `InvalidArgument`(?))
    * [ ] It could be reasonable to create a completed task, instead of first creating a non-completed task and then completing it with `SetCompleted`. That is a later use case though.

## `UpdateTask`

* [ ] With empty `name`: fail with `InvalidArgument`
* [ ] With `name` field with invalid format: fail with `InvalidArgument`
* [ ] With non-existent task: fail with `NotFound`
* [ ] Setting output only fields: shoud fail with `InvalidArgument`:
    * [ ] Both for empty `update_mask` (i.e., update all fields), or for `update_mask` specifically containing the output-only fields
* [ ] Updating the `title` to be empty

## `DeleteTask`

* [ ] Empty `name` field: fail with `InvalidArgument`
* [ ] Malformed `name` field: fail with `InvalidArgument`
* [ ] Non-existent task: fail with `NotFound`
* [ ] Deleting already deleted task: fail with `FailedPrecondition`

## `SetCompleted`

* [ ] Empty `name` field: fail with `InvalidArgument`
* [ ] Malformed `name` field: fail with `InvalidArgument`
* [ ] Non-existent task: fail with `NotFound`
* [ ] Not modifying (i.e., `completed = false -> completed = false` or `completed = true -> completed = true`):
    * [ ] Verify that `updated = false`
	* [ ] `GetTask` before and after, there should be no difference
* [ ] Modifying (i.e., `completed = false -> completed = true` or `completed = true -> completed = false`):
    * [ ] Verify that `updated = true`
	* [ ] `GetTask` before and after: should only diff in `completed` field
