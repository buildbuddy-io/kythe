#
# Copyright 2016 The Kythe Authors. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

# Bazel rules to extract Go compilations from library targets for testing the
# Go cross-reference indexer.
load("@bazel_skylib//lib:shell.bzl", "shell")
load(
    "@io_bazel_rules_go//go:def.bzl",
    "GoInfo",
    "go_library",
)
load(
    "//tools/build_rules/verifier_test:verifier_test.bzl",
    "KytheEntryProducerInfo",
    "kythe_integration_test",
    "verifier_test",
)

# Emit a shell script that sets up the environment needed by the extractor to
# capture dependencies and runs the extractor.
def _emit_extractor_script(ctx, mode, script, output, srcs, deps, ipath, data, extra_extractor_args):
    tmpdir = output.dirname + "/tmp"
    srcroot = tmpdir + "/src"
    srcdir = srcroot + "/" + ipath
    extras = []
    cmds = ["#!/bin/sh -e", "mkdir -p " + srcdir]

    # Link the source files and dependencies into a common temporary directory.
    # Source files need to be made relative to the temp directory.
    ups = srcdir.count("/") + 1
    cmds += [
        'ln -s "%s%s" "%s"' % ("../" * ups, src.path, srcdir)
        for src in srcs
    ]
    for dep in deps:
        gosrc = dep.source
        path = gosrc.importpath
        fullpath = "/".join([srcroot, path])
        tups = fullpath.count("/") + 1
        cmds += ["mkdir -p " + fullpath]
        for src in gosrc.srcs:
            cmds += ["ln -s '%s%s' '%s'" % ("../" * tups, src.path, fullpath + "/" + src.basename)]

    # Gather any extra data dependencies.
    for target in data:
        for f in target.files.to_list():
            cmds.append('ln -s "%s%s" "%s"' % ("../" * ups, f.path, srcdir))
            extras.append(srcdir + "/" + f.path.rsplit("/", 1)[-1])

    # Invoke the extractor on the temp directory.
    goroot = "/".join(ctx.files._sdk_files[0].path.split("/")[:-2])
    cmds.append("export GOCACHE=\"$PWD/" + tmpdir + "/cache\"")
    cmds.append("export CGO_ENABLED=0")

    args = [ctx.files._extractor[-1].path] + extra_extractor_args + [
        "-output",
        output.path,
        "-goos",
        mode.goos,
        "-goarch",
        mode.goarch,
        "-goroot",
        goroot,
        "-gocompiler",
        "gc",
        "-gopath",
        tmpdir,
        "-extra_files",
        "'%s'" % ",".join(extras),
        ipath,
    ]
    cmds.append(" ".join(args))

    f = ctx.actions.declare_file(script)
    ctx.actions.write(output = f, content = "\n".join(cmds), is_executable = True)
    return f

KZipProvider = provider()

def _go_extract(ctx):
    gosrc = ctx.attr.library[GoInfo]
    mode = gosrc.mode
    srcs = gosrc.srcs

    # TODO: handle transitive dependencies
    deps = gosrc.deps
    depsrcs = []
    for dep in deps:
        depsrcs += dep.source.srcs

    ipath = gosrc.importpath
    data = ctx.attr.data
    output = ctx.outputs.kzip
    script = _emit_extractor_script(
        ctx,
        mode,
        ctx.label.name + "-extract.sh",
        output,
        srcs,
        deps,
        ipath,
        data,
        ctx.attr.extra_extractor_args,
    )

    extras = []
    for target in data:
        extras += target.files.to_list()

    tools = ctx.files._extractor + ctx.files._sdk_files
    ctx.actions.run(
        mnemonic = "GoExtract",
        executable = script,
        outputs = [output],
        inputs = srcs + extras + depsrcs,
        tools = tools,
    )
    return [KZipProvider(kzip = output)]

# Generate a kzip with the compilations captured from a single Go library or
# binary rule.
go_extract = rule(
    _go_extract,
    attrs = {
        # Additional data files to include in each compilation.
        "data": attr.label_list(
            allow_empty = True,
            allow_files = True,
        ),
        "library": attr.label(
            providers = [GoInfo],
            mandatory = True,
        ),
        "_extractor": attr.label(
            default = Label("//kythe/go/extractors/cmd/gotool"),
            executable = True,
            cfg = "exec",
        ),
        "_sdk_files": attr.label(
            allow_files = True,
            default = "//third_party:go_sdk_files",
        ),
        "extra_extractor_args": attr.string_list(),
    },
    outputs = {"kzip": "%{name}.kzip"},
    toolchains = ["@io_bazel_rules_go//go:toolchain"],
)

def _go_entries(ctx):
    kzip = ctx.attr.kzip[KZipProvider].kzip
    iargs = []
    output = ctx.outputs.entries

    # If the test wants marked source, enable support for it in the indexer.
    if ctx.attr.has_marked_source:
        iargs.append("-code")

    if ctx.attr.emit_anchor_scopes:
        iargs.append("-anchor_scopes")

    if ctx.attr.use_compilation_corpus_for_all:
        iargs.append("-use_compilation_corpus_for_all")

    if ctx.attr.use_file_as_top_level_scope:
        iargs.append("-use_file_as_top_level_scope")

    if ctx.attr.override_stdlib_corpus:
        iargs.append("-override_stdlib_corpus=%s" % ctx.attr.override_stdlib_corpus)

    # If the test wants linkage metadata, enable support for it in the indexer.
    if ctx.attr.metadata_suffix:
        iargs += ["-meta", ctx.attr.metadata_suffix]

    test_runners = []
    eargs = [ctx.expand_location(arg.replace("$(location", "$(rootpath"), ctx.attr.data) for arg in ctx.attr.extra_indexer_args]
    test_runners.append(_make_test_runner(ctx, {}, arguments = iargs + eargs + [kzip.short_path]))

    iargs += [ctx.expand_location(arg, ctx.attr.data) for arg in ctx.attr.extra_indexer_args]
    iargs += [kzip.path, "| gzip >" + output.path]
    iargs.insert(0, ctx.executable._exec_indexer.path)

    cmds = ["set -e", "set -o pipefail", " ".join(iargs), ""]
    ctx.actions.run_shell(
        mnemonic = "GoIndexer",
        command = "\n".join(cmds),
        outputs = [output],
        inputs = [kzip] + ctx.files.data,
        tools = [ctx.executable._exec_indexer],
    )

    return [
        KytheEntryProducerInfo(
            executables = test_runners,
            runfiles = ctx.runfiles(
                files = (test_runners + [kzip] + ctx.files.data),
            ).merge(ctx.attr._indexer[DefaultInfo].default_runfiles),
        ),
    ]

# Run the Kythe indexer on the output that results from a go_extract rule.
go_entries = rule(
    _go_entries,
    attrs = {
        # Whether to enable explosion of MarkedSource facts.
        "has_marked_source": attr.bool(default = False),

        # Whether to enable anchor scope edges.
        "emit_anchor_scopes": attr.bool(default = False),

        # The go_extract output to pass to the indexer.
        "kzip": attr.label(
            providers = [KZipProvider],
            mandatory = True,
        ),

        # The suffix used to recognize linkage metadata files, if non-empty.
        "metadata_suffix": attr.string(default = ""),
        "use_compilation_corpus_for_all": attr.bool(default = False),
        "use_file_as_top_level_scope": attr.bool(default = False),
        "override_stdlib_corpus": attr.string(default = ""),
        "extra_indexer_args": attr.string_list(),

        # Extra files required by the indexer
        "data": attr.label_list(
            allow_empty = True,
            allow_files = True,
        ),

        # The location of the Go indexer binary.
        "_indexer": attr.label(
            default = Label("//kythe/go/indexer/cmd/go_indexer"),
            executable = True,
            cfg = "target",
        ),
        "_exec_indexer": attr.label(
            default = Label("//kythe/go/indexer/cmd/go_indexer"),
            executable = True,
            cfg = "exec",
        ),
        "_test_template": attr.label(
            default = Label("//tools/build_rules/verifier_test:indexer.sh.in"),
            allow_single_file = True,
        ),
    },
    outputs = {"entries": "%{name}.entries.gz"},
)

def _make_test_runner(ctx, env, arguments):
    output = ctx.actions.declare_file(ctx.label.name + "_test_runner")
    ctx.actions.expand_template(
        output = output,
        is_executable = True,
        template = ctx.file._test_template,
        substitutions = {
            "@INDEXER@": shell.quote(ctx.executable._indexer.short_path),
            "@ENV@": "\n".join([
                shell.quote("{key}={value}".format(key = key, value = value))
                for key, value in env.items()
            ]),
            "@ARGS@": "\n".join([
                shell.quote(a)
                for a in arguments
            ]),
        },
    )
    return output

def go_verifier_test(
        name,
        entries,
        srcs = [],
        deps = [],
        size = "small",
        tags = [],
        log_entries = False,
        has_marked_source = False,
        resolve_code_facts = False,
        allow_duplicates = False,
        use_fast_solver = False):
    opts = ["--use_file_nodes", "--show_goals", "--check_for_singletons", "--goal_regex='\\s*//\\s?-(.*)'"]
    if log_entries:
        opts.append("--show_protos")
    if allow_duplicates or len(deps) > 0:
        opts.append("--ignore_dups")
    if len(srcs) > 0:
        opts.append("--nofile_vnames")

    # If the test wants marked source, enable support for it in the verifier.
    if has_marked_source:
        opts.append("--convert_marked_source")
    if not use_fast_solver:
        opts.append("--use_fast_solver=false")
    return verifier_test(
        name = name,
        size = size,
        opts = opts,
        tags = tags,
        resolve_code_facts = resolve_code_facts,
        srcs = srcs,
        deps = [entries] + deps,
    )

# Shared extract/index logic for the go_indexer_test/go_integration_test rules.
def _go_indexer(
        name,
        srcs,
        deps = [],
        importpath = None,
        data = None,
        has_marked_source = False,
        emit_anchor_scopes = False,
        allow_duplicates = False,
        use_compilation_corpus_for_all = False,
        use_file_as_top_level_scope = False,
        override_stdlib_corpus = "",
        metadata_suffix = "",
        extra_indexer_args = [],
        extra_extractor_args = []):
    if importpath == None:
        importpath = native.package_name() + "/" + name
    lib = name + "_lib"
    go_library(
        name = lib,
        srcs = srcs,
        importpath = importpath,
        deps = [dep + "_lib" for dep in deps],
    )
    kzip = name + "_units"
    go_extract(
        name = kzip,
        data = data,
        library = lib,
        extra_extractor_args = extra_extractor_args,
    )
    entries = name + "_entries"
    go_entries(
        name = entries,
        data = data,
        has_marked_source = has_marked_source,
        emit_anchor_scopes = emit_anchor_scopes,
        use_compilation_corpus_for_all = use_compilation_corpus_for_all,
        use_file_as_top_level_scope = use_file_as_top_level_scope,
        override_stdlib_corpus = override_stdlib_corpus,
        extra_indexer_args = extra_indexer_args,
        kzip = ":" + kzip,
        metadata_suffix = metadata_suffix,
        tags = ["manual"],
    )
    return entries

# A convenience macro to generate a test library, pass it to the Go indexer,
# and feed the output of indexing to the Kythe schema verifier.
def go_indexer_test(
        name,
        srcs,
        deps = [],
        import_path = None,
        size = None,
        tags = None,
        log_entries = False,
        data = None,
        has_marked_source = False,
        resolve_code_facts = False,
        emit_anchor_scopes = False,
        allow_duplicates = False,
        use_compilation_corpus_for_all = False,
        use_file_as_top_level_scope = False,
        override_stdlib_corpus = "",
        metadata_suffix = "",
        extra_goals = [],
        extra_indexer_args = [],
        extra_extractor_args = [],
        use_fast_solver = False):
    entries = _go_indexer(
        name = name,
        srcs = srcs,
        data = data,
        has_marked_source = has_marked_source,
        emit_anchor_scopes = emit_anchor_scopes,
        use_compilation_corpus_for_all = use_compilation_corpus_for_all,
        use_file_as_top_level_scope = use_file_as_top_level_scope,
        override_stdlib_corpus = override_stdlib_corpus,
        importpath = import_path,
        metadata_suffix = metadata_suffix,
        deps = deps,
        extra_indexer_args = extra_indexer_args,
        extra_extractor_args = extra_extractor_args,
    )
    go_verifier_test(
        name = name,
        srcs = extra_goals,
        size = size,
        allow_duplicates = allow_duplicates,
        entries = ":" + entries,
        deps = [dep + "_entries" for dep in deps],
        has_marked_source = has_marked_source,
        resolve_code_facts = resolve_code_facts,
        log_entries = log_entries,
        tags = tags,
        use_fast_solver = use_fast_solver,
    )

# A convenience macro to generate a test library, pass it to the Go indexer,
# and feed the output of indexing to the Kythe integration test pipeline.
def go_integration_test(
        name,
        srcs,
        deps = [],
        data = None,
        file_tickets = [],
        import_path = None,
        size = "small",
        has_marked_source = False,
        metadata_suffix = ""):
    entries = _go_indexer(
        name = name,
        srcs = srcs,
        data = data,
        has_marked_source = has_marked_source,
        import_path = import_path,
        metadata_suffix = metadata_suffix,
        deps = deps,
    )
    kythe_integration_test(
        name = name,
        size = size,
        srcs = [":" + entries],
        file_tickets = file_tickets,
    )
