package testcase

import "github.com/bazelbuild/rules_go/go/tools/bazel"

func Runfile(path string) string {
	rf, err := bazel.Runfile(path)
	if err != nil {
		panic(err)
	}
	return rf
}
