# GoMW
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/jnnkrdb/gomw)](https://github.com/jnnkrdb/gomw)
[![CodeFactor](https://www.codefactor.io/repository/github/jnnkrdb/gomw/badge)](https://www.codefactor.io/repository/github/jnnkrdb/gomw)
[![Go Report Card](https://goreportcard.com/badge/github.com/jnnkrdb/gomw)](https://goreportcard.com/report/github.com/jnnkrdb/gomw)
[![Github tag](https://badgen.net/github/tag/jnnkrdb/gomw)](https://github.com/jnnkrdb/gomw/tags/)
[![GitHub issues](https://badgen.net/github/issues/jnnkrdb/gomw/)](https://github.com/jnnkrdb/gomw/issues/)
[![GPLv3 license](https://img.shields.io/badge/License-GPLv3-blue.svg)](http://perso.crans.org/besson/LICENSE.html)

---
## Description
A go-package for http-middleware chain supply.

## Install
Use the go commandline tool to install the package to your go project.
```
go get github.com/jnnkrdb/gomw
```

## How to use

This package translates the default middleware chain 
```go
var mux *http.ServeMux = http.NewServeMux()
mux.Handle("/pattern", middleware1(middleware2(middleware3(func(w http.ResponseWriter, r *http.Request)))))
http.ListenAndServe(":80", mux)
```
into
```go
var mux *http.ServeMux = http.NewServeMux()
mux.Handle("/pattern", middlewares.New(middleware1, middleware2, middleware3).Then(func(w http.ResponseWriter, r *http.Request)))
http.ListenAndServe(":80", mux)
```
or even easier, with the addition handlers package, for multiple patterns with function, you can translate the functions directly into the http.Handler for your serving.
```go
package main

import (
	"net/http"

	mw "github.com/jnnkrdb/gomw/middlewares"
	hndlrs "github.com/jnnkrdb/gomw/handlers"
)

func main() {
  var funcArr = hndlrs.HttpFunctionSet{
    hndlrs.HttpFunction{
      Pattern: "/function_A",
      Function: func(http.ResponseWriter, *http.Request) {
        w.Write([]byte("/function_A"))
      },
      MiddleWares: mw.MiddleWareChain{
        middleware1,
        middleware2,
        middleware3,
      },
    },
    hndlrs.HttpFunction{
      Pattern: "/function_B",
      Function: func(http.ResponseWriter, *http.Request) {
        w.Write([]byte("/function_B"))
      },
      MiddleWares: mw.MiddleWareChain{
        middleware_A,
        middleware_B,
        middleware_C,
      },

    },
    hndlrs.HttpFunction{
      Pattern: "/function_C",
      Function: func(http.ResponseWriter, *http.Request) {
        w.Write([]byte("/function_C"))
      },
      MiddleWares: mw.MiddleWareChain{
        middleware1,
        middleware_2,
        middleware_XY,
      },
    },
  }

  http.ListenAndServe(":80", hndlrs.GetHandler(funcArr))
}
```