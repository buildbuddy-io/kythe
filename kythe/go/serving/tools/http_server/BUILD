load("//tools:build_rules/shims.bzl", "go_binary")

package(default_visibility = ["//kythe:default_visibility"])

go_binary(
    name = "http_server",
    srcs = ["http_server.go"],
    deps = [
        "//kythe/go/services/filetree",
        "//kythe/go/services/graph",
        "//kythe/go/services/graphstore/proxy",
        "//kythe/go/services/xrefs",
        "//kythe/go/serving/filetree",
        "//kythe/go/serving/graph",
        "//kythe/go/serving/identifiers",
        "//kythe/go/serving/xrefs",
        "//kythe/go/storage/kvutil",
        "//kythe/go/storage/leveldb",
        "//kythe/go/storage/pebble",
        "//kythe/go/storage/table",
        "//kythe/go/util/flagutil",
        "//kythe/go/util/log",
        "@org_golang_x_net//http2",
    ],
)
