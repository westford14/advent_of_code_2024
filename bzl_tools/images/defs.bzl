"""Rules and macros for building and pushing images"""

load("@rules_oci//oci:defs.bzl", "oci_image", "oci_load", "oci_push")

def docker_image(name, entrypoint, tars, image_name, base = "@go_base"):
    """Macro for all the dodads needed for an image.

    This creates rules for building, loading and pushing images,
    and injects build-time labels and environment variables

    Args:
        name: name
        entrypoint: Entrypoint for your image
        tars: tars to be included in your image
        image_name: ECR repo name, should usually be the same as service_name
        base: Base image
    """
    imagetarget_name = name + "_image"
    oci_image(
        name = imagetarget_name,
        base = base,
        entrypoint = entrypoint,
        env = ":env",
        tars = tars,
        labels = ":labels",
    )
    oci_load(
        name = name + "_load",
        image = ":" + imagetarget_name,
        repo_tags = [image_name + ":latest"],
    )
    oci_push(
        name = name + "_push",
        image = ":" + imagetarget_name,
        remote_tags = ["bogus"],
        repository = "1234567890.dkr.ecr.us-west-2.amazonaws.com/" + image_name,
    )
