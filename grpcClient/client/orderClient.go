/*
@Time : 2020/5/15 17:43
@Author : zxr
@File : orderClient
@Software: GoLand
*/
package main

import (
	"bili/grpcClient/pbfiles"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

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
	client := pbfiles.NewOrderServiceClient(conn)
	ctx := context.Background()
	orderReq := &pbfiles.OrderNewRequest{
		OrderMain: &pbfiles.OrderMain{
			OrderId:    123,
			OrderNo:    "sn123",
			OrderMoney: 1.23,
			OrderTime:  &timestamp.Timestamp{Seconds: time.Now().Unix()},
		},
	}
	resp, e := client.NewOrder(ctx, orderReq)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println("NewOrder:", resp)
}
