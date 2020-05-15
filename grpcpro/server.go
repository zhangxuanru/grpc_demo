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
	"log"
	"net"

	"google.golang.org/grpc/credentials"

	"google.golang.org/grpc"
)

//grpc 服务端
func main() {
	//无证书 start
	//rpcServer := grpc.NewServer()
	//pbfiles.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	//listener, _ := net.Listen("tcp", ":8081")
	//if err := rpcServer.Serve(listener); err != nil {
	//	fmt.Println("error:", err)
	//}
	//无证书 end

	creds, e := credentials.NewServerTLSFromFile("cert/server.crt", "cert/server.key")
	if e != nil {
		log.Fatal(e)
	}
	rpcServer := grpc.NewServer(grpc.Creds(creds))
	pbfiles.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	pbfiles.RegisterOrderServiceServer(rpcServer, new(services.Order))
	listener, _ := net.Listen("tcp", ":8081")
	if e := rpcServer.Serve(listener); e != nil {
		log.Fatal(e)
	}
}

//加上证书验证
func main1() {
	creds, e := credentials.NewServerTLSFromFile("cert/server.crt", "cert/server.key")
	if e != nil {
		log.Fatal(e)
	}
	rpcServer := grpc.NewServer(grpc.Creds(creds))
	pbfiles.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	listener, _ := net.Listen("tcp", ":8081")
	if e := rpcServer.Serve(listener); e != nil {
		log.Fatal(e)
	}
}
