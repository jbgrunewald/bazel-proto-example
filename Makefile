update-repos:
	bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro build_tools/go_deps.bzl%go_deps