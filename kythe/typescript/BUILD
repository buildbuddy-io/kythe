load("@aspect_rules_jasmine//jasmine:defs.bzl", "jasmine_test")
load("@aspect_rules_js//js:defs.bzl", "js_test")
load("@aspect_rules_ts//ts:defs.bzl", "ts_config", "ts_project")

package(default_visibility = ["//kythe:default_visibility"])

# If you need to add / update dependencies:
#
# 0. Ensure you have a recent npm and pnpm. npm can come from apt.
# 1. Edit //:package.json and //:WORKSPACE as appropriate
# 2. From the *root dir* of the project, run
#
#       npm install
#       pnpm import
#
# 3. Add the dependencies to the ts_project build rule below.

ts_project(
    name = "kythe",
    srcs = ["kythe.ts"],
    declaration = True,
    source_map = True,
    supports_workers = 0,
)

ts_project(
    name = "indexer",
    srcs = [
        "indexer.ts",
        "plugin_api.ts",
        "utf8.ts",
    ],
    declaration = True,
    source_map = True,
    supports_workers = 0,
    deps = [
        ":kythe",
        "//:node_modules/@types/node",
        "//:node_modules/typescript",
    ],
)

ts_project(
    name = "test_ts",
    srcs = ["test.ts"],
    declaration = True,
    source_map = True,
    supports_workers = 0,
    deps = [
        ":indexer",
        ":kythe",
        "//:node_modules/@types/node",
        "//:node_modules/source-map",
        "//:node_modules/source-map-support",
        "//:node_modules/typescript",
    ],
)

ts_project(
    name = "utf8_test_ts",
    testonly = True,
    srcs = ["utf8_test.ts"],
    declaration = True,
    source_map = True,
    supports_workers = 0,
    deps = [
        ":indexer",
        "//:node_modules/@types/jasmine",
        "//:node_modules/balanced-match",
        "//:node_modules/brace-expansion",
        "//:node_modules/fs.realpath",
        "//:node_modules/glob",
        "//:node_modules/inflight",
        "//:node_modules/inherits",
        "//:node_modules/jasmine",
        "//:node_modules/jasmine-core",
        "//:node_modules/jasmine-reporters",
        "//:node_modules/minimatch",
        "//:node_modules/once",
        "//:node_modules/typescript",
        "//:node_modules/wrappy",
    ],
)

js_test(
    name = "indexer_test",
    data = [
        ":test_ts",
    ] + glob(
        ["testdata/**/*"],
    ) + [
        "//kythe/cxx/verifier",
        "//kythe/go/platform/tools/entrystream",
        "//kythe/go/util/tools/markedsource",
    ],
    entry_point = ":test.js",
)

jasmine_test(
    name = "utf8_test",
    args = ["utf8_test.js"],
    chdir = package_name(),
    data = [
        "utf8_test.js",
        ":utf8_test_ts",
    ],
    include_transitive_sources = True,
    log_level = "debug",
    node_modules = "//:node_modules",
    patch_node_fs = True,
)
