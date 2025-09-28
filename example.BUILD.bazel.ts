load("@build_bazel_rules_typescript//:defs.bzl", "ts_library")
load("@rules_nodejs//:index.bzl", "nodejs_binary")  # or load proper rules_nodejs entrypoint

ts_library(
    name = "lib",
    srcs = glob(["*.ts"]),
    tsconfig = "//:tsconfig.json",
)

nodejs_binary(
    name = "app",
    data = [":lib"],
    entry_point = "node/app/main.ts",
)
