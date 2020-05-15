/*
@Time : 2020/5/13 18:50
@Author : zxr
@File : client
@Software: GoLand
*/
package main

import (
	"bili/grpcClient/pbfiles"
	"context"
	"fmt"

	"log"

	"google.golang.org/grpc/credentials"

	_ "google.golang.org/grpc/credentials"

	"google.golang.org/grpc"
)

//grpc client
func main() {
	creds, e := credentials.NewClientTLSFromFile("cert/server.crt", "bj")
	if e != nil {
		log.Fatal(e)
	}
	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(creds)) //grpc.WithInsecure()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pbfiles.NewProdServiceClient(conn)
	response, err := client.GetProdStock(context.Background(), &pbfiles.ProdRequest{ProdId: 10, Area: pbfiles.ProdArea_C})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)

	listResponse, e := client.GetProdList(context.Background(), &pbfiles.ProdListRequest{Size: 10})
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(listResponse)

	resp, err := client.GetProdInfo(context.Background(), &pbfiles.ProdRequest{ProdId: 10})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("resp:", resp)

	fmt.Println("----------------------- order:")
}
