package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/thinkgos/cdpui"
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
	go func() {
		time.Sleep(time.Second * 15)
		ui.Close()
	}()
	<-ui.Wait()
}

const indexHTML = `<!doctype html>
<html>
<head>
  <title>example</title>
</head>
<body>
  <div id="box3">
	<input id="input1" type="submit" value="dian">
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
