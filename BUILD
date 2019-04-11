load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")
load("@bazel_gazelle//:def.bzl", "gazelle")

gazelle(
    name = "gazelle",
    external = "vendored",
    prefix = "github.com/seibert-media/eventmeta",
)

go_library(
    name = "go_default_library",
    srcs = [
        "accessor.go",
        "context.go",
        "event.go",
        "group.go",
        "incomplete.go",
        "meta.go",
    ],
    importpath = "github.com/seibert-media/eventmeta",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/pkg/errors:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "incomplete_test.go",
        "meta_test.go",
    ],
    embed = [":go_default_library"],
    deps = [
        "//vendor/github.com/google/gofuzz:go_default_library",
        "//vendor/k8s.io/utils/diff:go_default_library",
    ],
)
