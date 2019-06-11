package(default_visibility = ["//kythe:default_visibility"])

load("@npm_bazel_typescript//:index.bzl", "ts_library", "ts_config")
load("@npm_bazel_jasmine//:index.bzl", "jasmine_node_test")

ts_library(
    name = "indexer",
    srcs = [
        "indexer.ts",
        "utf8.ts",
    ],
    tsconfig = ":tsconfig",
    deps = [
        "@npm//@types/node",
        "@npm//source-map-support",
        "@npm//typescript",
    ],
)

ts_library(
    name = "test_lib",
    srcs = [
        "test.ts",
        "utf8_test.ts",
    ],
    tsconfig = ":test_tsconfig",
    deps = [
        ":indexer",
        "@npm//@types/jasmine",
        "@npm//@types/node",
        "@npm//jasmine",
        "@npm//typescript",
    ],
)

jasmine_node_test(
    name = "indexer_test",
    srcs = [":test_lib"],
    data = [
        "//kythe/typescript/testdata",
    ],
)

ts_config(
    name = "tsconfig",
    src = "tsconfig.json",
    deps = [],
)

ts_config(
    name = "test_tsconfig",
    src = "tsconfig-test.json",
    deps = [":tsconfig"],
)
