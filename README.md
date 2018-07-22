# go-ptr

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/suzuki-shunsuke/go-ptr)
[![CircleCI](https://circleci.com/gh/suzuki-shunsuke/go-ptr.svg?style=svg)](https://circleci.com/gh/suzuki-shunsuke/go-ptr)
[![codecov](https://codecov.io/gh/suzuki-shunsuke/go-ptr/branch/master/graph/badge.svg)](https://codecov.io/gh/suzuki-shunsuke/go-ptr)
[![Go Report Card](https://goreportcard.com/badge/github.com/suzuki-shunsuke/go-ptr)](https://goreportcard.com/report/github.com/suzuki-shunsuke/go-ptr)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/go-ptr.svg)](https://github.com/suzuki-shunsuke/go-ptr)
[![GitHub tag](https://img.shields.io/github/tag/suzuki-shunsuke/go-ptr.svg)](https://github.com/suzuki-shunsuke/go-ptr/releases)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/go-ptr/master/LICENSE)

golang library to convert a literal to pointer.

To be honest, this library is very trivial.
If there are any other ways to get a pointer of the literal value, please tell us.

## Motivation

Sometimes we want to get a pointer of the literal value.
But the following code raises a compile error.

https://play.golang.org/p/p9E9omKHogv

```golang
package main

import (
	"fmt"
)

func main() {
	fmt.Println(&"hello")
}
```

```
prog.go:8:14: cannot take the address of "hello"
```

So we have to assign the literal to a variable.

```golang
a := "hello"
fmt.Println(&a)
```

But we want to write this at an one liner.

```golang
fmt.Println(ptr.PStr("hello"))
```

## License

[MIT](LICENSE)
