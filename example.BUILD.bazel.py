load("@rules_python//python:defs.bzl", "py_library", "py_binary")

py_library(
    name = "lib",
    srcs = ["lib.py"],
    visibility = ["//visibility:public"],
)

py_binary(
    name = "app",
    srcs = ["main.py"],
    deps = [":lib"],
)
