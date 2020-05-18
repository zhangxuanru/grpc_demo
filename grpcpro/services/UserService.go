/*
@Time : 2020/5/18 15:28
@Author : zxr
@File : UserService
@Software: GoLand
*/
package services

import (
	"bili/grpcpro/pbfiles"
	"context"
	"io"
)

type User struct {
}

//普通方法
func (u *User) GetUserScore(ctx context.Context, req *pbfiles.UserScoreRequest) (*pbfiles.UserScoreResponse, error) {
	var score int32 = 101
	users := make([]*pbfiles.UserInfo, 0)
	for _, user := range req.Users {
		user.UserScore = score
		score++
		users = append(users, user)
	}
	return &pbfiles.UserScoreResponse{Users: users}, nil
}

//服务端流
func (u *User) GetUserScoreByServerStream(req *pbfiles.UserScoreRequest, stream pbfiles.UserService_GetUserScoreByServerStreamServer) error {
	var score int32 = 101
	users := make([]*pbfiles.UserInfo, 0)
	for _, user := range req.Users {
		user.UserScore = score
		score++
		users = append(users, user)
		if len(users) >= 2 {
			if err := stream.Send(&pbfiles.UserScoreResponse{Users: users}); err != nil {
				return err
			}
			users = users[0:0]
		}
	}
	if len(users) > 0 {
		if err := stream.Send(&pbfiles.UserScoreResponse{Users: users}); err != nil {
			return err
		}
	}
	return nil
}

//客户端流
func (u *User) GetUserScoreByClientStream(stream pbfiles.UserService_GetUserScoreByClientStreamServer) error {
	var score int32 = 101
	users := make([]*pbfiles.UserInfo, 0)
	for {
		request, e := stream.Recv()
		if e == io.EOF {
			return stream.SendAndClose(&pbfiles.UserScoreResponse{Users: users})
		}
		if e != nil {
			return e
		}
		for _, user := range request.Users {
			user.UserScore = score
			score++
			users = append(users, user)
		}
	}
	return nil
}

//双向流
func (u *User) GetUserScoreTNS(stream pbfiles.UserService_GetUserScoreTNSServer) error {
	var score int32 = 101
	for {
		request, e := stream.Recv()
		if e == io.EOF {
			return nil
		}
		if e != nil {
			return e
		}
		users := make([]*pbfiles.UserInfo, 0)
		for _, user := range request.Users {
			user.UserScore = score
			score++
			users = append(users, user)
		}
		if e := stream.Send(&pbfiles.UserScoreResponse{Users: users}); e != nil {
			return e
		}
	}
	return nil
}
