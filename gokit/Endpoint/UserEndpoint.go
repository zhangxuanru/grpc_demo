/*
@Time : 2020/5/18 19:17
@Author : zxr
@File : UserEndpoint
@Software: GoLand
*/
package Endpoint

import (
	"bili/gokit/Service"
	"context"

	"github.com/go-kit/kit/endpoint"
)

type UserRequest struct {
	Uid int `json:"uid"`
}

type UserResponse struct {
	Result string `json:"result"`
}

func GenUserEndPoint(userService Service.IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(UserRequest)
		name := userService.GetName(r.Uid)
		return UserResponse{Result: name}, nil
	}
}
