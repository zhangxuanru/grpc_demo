/*
@Time : 2020/5/19 15:40
@Author : zxr
@File : main
@Software: GoLand
*/
package main

import (
	"bili/gokit/Endpoint"
	"bili/gokit/Service"
	"bili/gokit/Transport"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func main() {
	service := &Service.UserService{}
	point := Endpoint.GenUserEndPoint(service)
	handler := httpTransport.NewServer(point, Transport.DecodeUserRequest, Transport.EncodeUserResponse)

	r := mux.NewRouter()
	r.Handle(`/user/{uid:\d+}`, handler)
	//r.Methods(http.MethodGet).Path(`/user/{uid:\d+}`).Handler(handler)

	errChan := make(chan error)
	go func() {
		if err := http.ListenAndServe(":8080", r); err != nil {
			errChan <- err
		}
	}()

	go func() {
		sigChan := make(chan os.Signal)
		signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)
		errChan <- fmt.Errorf("signal:%s", <-sigChan)
	}()
	fmt.Println(<-errChan)
}
