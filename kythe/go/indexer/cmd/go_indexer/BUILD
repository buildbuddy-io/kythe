load("//tools:build_rules/shims.bzl", "go_binary")

package(default_visibility = ["//kythe:default_visibility"])

go_binary(
    name = "go_indexer",
    srcs = ["go_indexer.go"],
    deps = [
        "//kythe/go/indexer",
        "//kythe/go/platform/delimited",
        "//kythe/go/platform/kzip",
        "//kythe/go/platform/vfs",
        "//kythe/go/util/log",
        "//kythe/go/util/metadata",
        "//kythe/proto:analysis_go_proto",
        "//kythe/proto:go_go_proto",
        "//kythe/proto:storage_go_proto",
        "@com_github_golang_protobuf//proto",
        "@com_github_golang_protobuf//protoc-gen-go/descriptor",
        "@org_golang_google_protobuf//encoding/prototext",
    ],
)
