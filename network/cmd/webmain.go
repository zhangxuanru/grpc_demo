/*
@Time : 2020/5/21 16:02
@Author : zxr
@File : webmain
@Software: GoLand
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type web1 struct{}
type web2 struct{}
type proxy struct{}

func (w *web1) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("<h1>hello web1</h1>" + request.RemoteAddr))
}
func (w *web2) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("<h1>hello web2</h1>" + request.RemoteAddr))
}

//访问代理
func (w *proxy) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var r *http.Request
	webMap := map[int]string{
		9090: "http://localhost:9090",
		9091: "http://localhost:9091",
	}
	if request.URL.Path == "/a" {
		r, _ = http.NewRequest(request.Method, webMap[9090], request.Body)
	}
	if request.URL.Path == "/b" {
		r, _ = http.NewRequest(request.Method, webMap[9091], request.Body)
	}
	if r == nil {
		_, _ = writer.Write([]byte("<h1>hi zxr</h1>"))
		return
	}
	for k, v := range request.Header {
		r.Header.Set(k, v[0])
	}
	resp, _ := http.DefaultClient.Do(r)
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	for k, v := range resp.Header {
		writer.Header().Set(k, v[0])
	}
	writer.WriteHeader(resp.StatusCode)
	content, _ := ioutil.ReadAll(resp.Body)
	_, _ = fmt.Fprintf(writer, string(content))
}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	go func() {
		http.ListenAndServe(":8080", &proxy{})
	}()

	go func() {
		http.ListenAndServe(":9090", &web1{})
	}()

	go func() {
		http.ListenAndServe(":9091", &web2{})
	}()
	fmt.Println(<-c)
}

func testTrport() {
	//var r *http.Request
	////http.DefaultTransport.RoundTrip(r)
	//transport := http.Transport{
	//	ResponseHeaderTimeout: 1 * time.Second,
	//}
	//transport.RoundTrip(r)
	//parse, _ := url.Parse()
	////内置反向代理
	//reverseProxy := httputil.NewSingleHostReverseProxy(parse)
	//reverseProxy.ServeHTTP(w, r)
}
