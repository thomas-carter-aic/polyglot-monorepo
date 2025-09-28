load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_binary")

go_library(
    name = "go_default_library",
    srcs = ["hello.go"],
    importpath = "example.com/monorepo/go/hello",
    visibility = ["//visibility:public"],
)

go_binary(
    name = "hello_bin",
    srcs = ["main.go"],
    deps = [":go_default_library"],
)
# This BUILD.bazel file defines a Go library and a Go binary target.