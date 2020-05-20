/*
@Time : 2020/5/20 15:41
@Author : zxr
@File : http_rate
@Software: GoLand
*/
package main

import (
	"log"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

// http 限流，每次放一个， 总共5个令牌
var r = rate.NewLimiter(1, 5)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("OK"))
	})
	if err := http.ListenAndServe(":8080", rateLimit(mux)); err != nil {
		log.Fatal(err)
	}
}

//限流
func rateLimit(mux *http.ServeMux) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		//访问一次，取3个令牌
		if !r.AllowN(time.Now(), 3) {
			http.Error(writer, "too man connect....", http.StatusInternalServerError)
			return
		}
		mux.ServeHTTP(writer, request)
	})
}
