# pdp

A personal monorepo, ish. I try to collect all my projects here and build them all using Bazel.

## Setting up a local Kubernetes cluster

The `tools.go` file includes some tools used to maintain and build the code in
this repository. It includes the
[`sigs.k8s.io/kind`](https://pkg.go.dev/sigs.k8s.io/kind) package. This can be
run either using the `go` command or using Bazel.

*   With `go`, run the following (assuming you have `$GOPATH/bin` in `$PATH`):

    ```
    go install sigs.k8s.io/kind
    ```

    To create a cluster, run:

    ```
    kind create cluster [--name=foobar]
    ```

*   With Bazel, run:

    ```
    bazel run @io_k8s_sigs_kind//:kind -- create cluster [--name=foobar]
    ```