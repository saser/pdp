# pdp

A personal monorepo, ish. I try to collect all my projects here and build them all using Bazel.

## Setting up a local Kubernetes cluster with a local registry

[`kind`](https://kind.sigs.k8s.io) is used to set up a local Kubernetes cluster
for development. To be used effectively we also need a local registry where
Docker images can be pushed by the `rules_k8s` rules. The
`kind_with_registry.sh` script, run using the `//:kind_with_registry` target,
sets that up for us. So, simply doing this will get us started:

```
bazel run //:kind_with_registry
```

After this, BUILD packages can use the `k8s_local_object` repository to push
images to the created local repository.