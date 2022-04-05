# work
multi-module workspaces

```
$ pwd
/Users/ljh/Documents/work
$ 
$ find . -print | sed -e 's;[^/]*/;|____;g;s;____|; |;g' 
.
|____go.work
|____foo
| |____go.mod
| |____go.sum
| |____foo.go
| |____bar
| | |____bar.go
|____main
| |____go.mod
| |____go.sum
| |____main.go
$ 
$ 
$ cat go.work 
go 1.18

use (
	./foo
	./main
)
$ 
$ 
$ cat foo/go.mod 
module foo

go 1.18

require github.com/golang/glog v1.0.0
$ 
$ 
$ cat foo/go.sum 
github.com/golang/glog v1.0.0 h1:nfP3RFugxnNRyKgeWd4oI1nYvXpxrx8ck8ZrcizshdQ=
github.com/golang/glog v1.0.0/go.mod h1:EWib/APOK0SL3dFbYqvxE3UYd8E6s1ouQ7iEp/0LWV4=
$ 
$ 
$ cat foo/foo.go 
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
$ 
$ 
$ cat foo/bar/bar.go 
package bar

import (
	"github.com/golang/glog"
)

func Bar() {
	defer glog.Flush()
	glog.Info()
}
$ 
$ 
$ cat main/go.mod 
module main

go 1.18

require github.com/golang/glog v1.0.0
$ 
$   
$ cat main/go.sum 
github.com/golang/glog v1.0.0 h1:nfP3RFugxnNRyKgeWd4oI1nYvXpxrx8ck8ZrcizshdQ=
github.com/golang/glog v1.0.0/go.mod h1:EWib/APOK0SL3dFbYqvxE3UYd8E6s1ouQ7iEp/0LWV4=
$ 
$ 
$ cat main/main.go 
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
$ 

```
