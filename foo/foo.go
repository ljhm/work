package foo

import (
	"foo/bar"
	"github.com/golang/glog"
)

func Foo() {
	defer glog.Flush()
	glog.Info()
	bar.Bar()
}
