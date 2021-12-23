"""Contains a BUILD macro for creating Docker images of Postgres with a specific
schema file."""

load("@io_bazel_rules_docker//container:image.bzl", "container_image")
load("@rules_pkg//:pkg.bzl", "pkg_tar")

def postgres_image(
        name,
        schema = None):
    """Defines a container_image target with the given schema.

    Args:
        name: string. Required.
        schema: label. Required.
        """
    schema_tar = "_" + name + "_schema_tar"
    pkg_tar(
        name = schema_tar,
        srcs = [schema],
        package_dir = "docker-entrypoint-initdb.d",
    )
    container_image(
        name = name,
        base = "@postgres_image//image",
        env = {
            "POSTGRES_PASSWORD": name + "_password",
            "POSTGRES_USER": name + "_user",
        },
        tars = [schema_tar],
    )
