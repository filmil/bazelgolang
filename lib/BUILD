package(default_visibility=["//visibility:public"])
load("/tools/build_rules/go/def", "go_library")
load("/tools/build_rules/go/def", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["file.go"]
)

filegroup(
    name = "filegroup",
    srcs = ["file.txt"],
)
go_test(
    name = "go_default_library_test",
    srcs = ["file_test.go"],
    library = ":go_default_library",
    data = [":filegroup"],
)
