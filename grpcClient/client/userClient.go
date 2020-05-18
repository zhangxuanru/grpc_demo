/*
@Time : 2020/5/18 15:41
@Author : zxr
@File : userClient
@Software: GoLand
*/
package main

import (
	"bili/grpcClient/pbfiles"
	"context"
	"fmt"
	"io"
	"log"

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
	client := pbfiles.NewUserServiceClient(conn)

	//GetUserScoreByStream(client)
	//GetUserScoreByClientStream(client)
	GetUserScoreTNS(client)
}

//普通形式RPC调用
func GetUserScore(client pbfiles.UserServiceClient) {
	ctx := context.Background()
	var i int32
	req := &pbfiles.UserScoreRequest{}
	req.Users = make([]*pbfiles.UserInfo, 0)
	for i = 1; i < 10; i++ {
		req.Users = append(req.Users, &pbfiles.UserInfo{UserId: i})
	}
	if response, e := client.GetUserScore(ctx, req); e != nil {
		log.Fatal(e)
	} else {
		fmt.Println("response:", response)
	}
}

//服务端流形式RPC调用
func GetUserScoreByStream(client pbfiles.UserServiceClient) {
	end := make(chan bool)
	ctx := context.Background()
	var i int32
	req := &pbfiles.UserScoreRequest{}
	req.Users = make([]*pbfiles.UserInfo, 0)
	for i = 1; i < 10; i++ {
		req.Users = append(req.Users, &pbfiles.UserInfo{UserId: i})
	}
	streamClient, err := client.GetUserScoreByServerStream(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			response, err := streamClient.Recv()
			if err == io.EOF {
				end <- true
				break
			}
			if err != nil {
				end <- true
				log.Fatal("recv err:", err)
			}
			fmt.Println(response.Users)
		}
	}()
	<-end
}

//客户端流形式RPC调用
func GetUserScoreByClientStream(client pbfiles.UserServiceClient) {
	var i int32
	req := &pbfiles.UserScoreRequest{}
	ctx := context.Background()
	streamClient, err := client.GetUserScoreByClientStream(ctx)
	if err != nil {
		log.Fatal("err:", err)
	}
	for j := 1; j < 4; j++ {
		req.Users = make([]*pbfiles.UserInfo, 0)
		for i = 1; i < 6; i++ {
			req.Users = append(req.Users, &pbfiles.UserInfo{UserId: i})
		}
		if err := streamClient.Send(req); err != nil {
			log.Println("send err:", err)
		}
	}
	if response, err := streamClient.CloseAndRecv(); err != nil {
		log.Fatal("CloseAndRecv err:", err)
	} else {
		fmt.Println("response:", response)
	}
}

//双向流形式RPC调用
func GetUserScoreTNS(client pbfiles.UserServiceClient) {
	var i int32
	ctx := context.Background()
	req := &pbfiles.UserScoreRequest{}
	tnsClient, err := client.GetUserScoreTNS(ctx)
	if err != nil {
		log.Fatal("err:", err)
	}
	for j := 1; j < 4; j++ {
		req.Users = make([]*pbfiles.UserInfo, 0)
		for i = 1; i < 4; i++ {
			req.Users = append(req.Users, &pbfiles.UserInfo{UserId: i})
		}
		if err := tnsClient.Send(req); err != nil {
			log.Println("send err:", err)
		}
		response, err := tnsClient.Recv()
		if err != nil {
			log.Println("Recv err:", err)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("err:", err)
		}
		fmt.Println("Recv response:", response)
	}
}
