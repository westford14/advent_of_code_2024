"""
Rules for copying generated code back to source.
This allows devs to use the native go toolchain to dev,
and makes for a much better editor experience (since the
editor can only see the files in source)
"""

load("@bazel_skylib//lib:shell.bzl", "shell")

def _gen_copy_files_script(ctx, files, prefix = "github.com/westford14/advent_of_code_2024"):
    content = ""
    for f in files:
        idx = f.path.index(prefix) + len(prefix)
        source_path = f.path[idx:].lstrip("/")
        line = "mkdir -p $(dirname %s) && cp -f %s %s;\n" % (source_path, f.path, source_path)
        content += line
    substitutions = {
        "@@CONTENT@@": shell.quote(content),
    }
    out = ctx.actions.declare_file(ctx.label.name + ".sh")
    ctx.actions.expand_template(
        template = ctx.file._template,
        output = out,
        substitutions = substitutions,
        is_executable = True,
    )
    runfiles = ctx.runfiles(files = files)
    return [
        DefaultInfo(
            files = depset([out]),
            runfiles = runfiles,
            executable = out,
        ),
    ]

def _go_proto_link_impl(ctx):
    return _gen_copy_files_script(ctx, ctx.attr.dep[OutputGroupInfo].go_generated_srcs.to_list())

_go_proto_link = rule(
    implementation = _go_proto_link_impl,
    attrs = {
        "dir": attr.string(),
        "dep": attr.label(),
        "_template": attr.label(
            default = "//bzl_tools/protogen_go_link:copy_into_workspace.sh",
            allow_single_file = True,
        ),
    },
)

def go_proto_link(name, **kwargs):
    """Wrapper for go_proto_link_impl

    Args:
        name: name
        **kwargs: keyword args
    """
    if not "dir" in kwargs:
        dir = native.package_name()
        kwargs["dir"] = dir

    gen_rule_name = "%s_copy_gen" % name
    _go_proto_link(name = gen_rule_name, **kwargs)
    native.sh_binary(
        name = name,
        srcs = [":%s" % gen_rule_name],
    )
