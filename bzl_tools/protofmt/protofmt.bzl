"""
Formats protobuf files
"""

_TOOLCHAIN = "@rules_buf//tools/buf:toolchain_type"

def _buf_format(ctx):
    buf = ctx.toolchains[_TOOLCHAIN].cli

    ctx.actions.write(
        output = ctx.outputs.executable,
        content = """
{} format $BUILD_WORKSPACE_DIRECTORY/proto -w
""".format(buf.short_path),
        is_executable = True,
    )

    files = [buf]
    runfiles = ctx.runfiles(
        files = files,
    )

    return [
        DefaultInfo(
            runfiles = runfiles,
        ),
    ]

buf_format = rule(
    implementation = _buf_format,
    toolchains = [_TOOLCHAIN],
    executable = True,
)
