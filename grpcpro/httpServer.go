/*
@Time : 2020/5/13 18:34
@Author : zxr
@File : server
@Software: GoLand
*/
package main

import (
	"bili/grpcpro/pbfiles"
	"bili/grpcpro/services"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc/credentials"

	"google.golang.org/grpc"
)

//grpc http 服务端
func main() {
	rpcGatWayHttp()
}

//http2的方式提供服务端
func http2Server() {
	creds, e := credentials.NewServerTLSFromFile("cert/server.crt", "cert/server.key")
	if e != nil {
		log.Fatal(e)
	}
	rpcServer := grpc.NewServer(grpc.Creds(creds))
	pbfiles.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request)
		rpcServer.ServeHTTP(writer, request)
	})
	server := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}
	if err := server.ListenAndServeTLS("cert/server.crt", "cert/server.key"); err != nil {
		log.Fatal(err)
	}
}

//RPC 与HTTP的方式 同时提供服务
func rpcGatWayHttp() {
	rpcHost := "localhost:8081"
	mux := runtime.NewServeMux()
	creds, e := credentials.NewClientTLSFromFile("cert/server.crt", "bj")
	if e != nil {
		log.Fatal(e)
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	err := pbfiles.RegisterProdServiceHandlerFromEndpoint(context.Background(), mux, rpcHost, opts)
	err = pbfiles.RegisterOrderServiceHandlerFromEndpoint(context.Background(), mux, rpcHost, opts)
	if err != nil {
		log.Fatal(err)
	}
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	http.ListenAndServe(":8080", mux)
}
