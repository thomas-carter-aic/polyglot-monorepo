load("@bazel_tools//tools/jdk:default_java_toolchain.bzl", "default_java_toolchain")
# Use the `java_library` rule (no external-dep example here)
java_library(
    name = "mylib",
    srcs = glob(["src/main/java/**/*.java"]),
    visibility = ["//visibility:public"],
)
