package main

import (
	"flag"
	"foo"
	"foo/bar"
	"github.com/golang/glog"
)

func init() {
	flag.Parse()
	flag.Set("alsologtostderr", "true")
	glog.Info("-log_dir=", flag.Lookup("log_dir").Value)
}

func main() {
	defer glog.Flush()
	glog.Info()
	foo.Foo()
	glog.Info()
	bar.Bar()
}
