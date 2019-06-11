package(default_visibility = ["//kythe:default_visibility"])

load("@npm_bazel_typescript//:index.bzl", "ts_config")

ts_config(
    name = "tsconfig",
    src = "tsconfig.json",
    deps = [],
)

filegroup(
    name = "testdata",
    srcs = glob(["**/*"]),
)
