load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["encoding.go"],
    importmap = "github.com/kubernetes-sigs/poseidon/vendor/golang.org/x/text/encoding",
    importpath = "golang.org/x/text/encoding",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/text/encoding/internal/identifier:go_default_library",
        "//vendor/golang.org/x/text/transform:go_default_library",
    ],
)
