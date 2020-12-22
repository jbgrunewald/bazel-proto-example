load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")
load("@io_bazel_rules_go//proto:compiler.bzl", "go_proto_compiler")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")
load("@rules_proto//proto:defs.bzl", "proto_library")

package(
    default_visibility = ["//visibility:public"],
)

proto_library(
    name = "config_proto",
    srcs = [
        "config.proto",
    ],
    deps = ["@com_google_protobuf//:struct_proto"],
)

go_proto_compiler(
    name = "protoc-gen-go-json",
    plugin = "@com_github_mitchellh_protoc_gen_go_json//:protoc-gen-go-json",
    suffix = ".pb.json.go",
    valid_archive = False,
    deps = [
        "@com_github_golang_protobuf//jsonpb:go_default_library"
    ]
)

go_proto_library(
    name = "config_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_proto", ":protoc-gen-go-json"],
    importpath = "github.com/jbgrunewald/bazel-proto-example/proto/config",
    proto = ":config_proto",
)

go_library(
    name = "go_default_library",
    embed = [":config_go_proto"],
    importpath = "github.com/jbgrunewald/bazel-proto-example/proto/config",
)

filegroup(
    name = "proto_sources",
    srcs = glob(["*.proto"]),
)

go_test(
    name = "go_default_test",
    srcs = [
        "json_serialization_test.go",
    ],
    embed = [":go_default_library"]
)
