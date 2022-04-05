package bar

import (
	"github.com/golang/glog"
)

func Bar() {
	defer glog.Flush()
	glog.Info()
}
