# cdpui

only support headless_shell,chrome,chromium

[![GoDoc](https://godoc.org/github.com/things-go/cdpui?status.svg)](https://godoc.org/github.com/things-go/cdpui)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/things-go/cdpui?tab=doc)
[![codecov](https://codecov.io/gh/things-go/cdpui/branch/main/graph/badge.svg)](https://codecov.io/gh/things-go/cdpui)
![Action Status](https://github.com/things-go/cdpui/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/things-go/cdpui)](https://goreportcard.com/report/github.com/things-go/cdpui)
[![License](https://img.shields.io/github/license/things-go/cdpui)](https://github.com/things-go/cdpui/raw/main/LICENSE)
[![Tag](https://img.shields.io/github/v/tag/things-go/cdpui)](https://github.com/things-go/cdpui/tags)

## example 

[embedmd]:# (_example/main.go go)
```go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/things-go/cdpui"
)

func main() {
	ui := cdpui.New("http://localhost:8080")
	defer ui.Close()

	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
			fmt.Fprint(w, indexHTML)
		})
		http.ListenAndServe(":8080", nil)
	}()
	time.Sleep(time.Millisecond * 5)
	ui.Run()
	// go func() {
	// 	time.Sleep(time.Second * 15)
	// 	ui.Close()
	// }()
	<-ui.Wait()
}

const indexHTML = `<!doctype html>
<html>
<head>
  <title>example</title>
</head>
<body>
  <div>
	<input id="input1" type="submit" value="dian">
  </div>
  <div id="box3">
	<textarea id="textarea1" style="width:500px;height:400px">textarea</textarea><br><br>
	<input id="input2" type="submit" value="Next">
	<select id="select1">
		<option value="one">1</option>
		<option value="two">2</option>
		<option value="three">3</option>
		<option value="four">4</option>
	</select>
  </div>
</body>
</html>`
```
