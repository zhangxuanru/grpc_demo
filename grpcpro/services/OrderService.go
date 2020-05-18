/*
@Time : 2020/5/15 17:21
@Author : zxr
@File : OrderService
@Software: GoLand
*/
package services

import (
	. "bili/grpcpro/pbfiles"
	"context"
	"fmt"
)

type Order struct {
}

func (o *Order) NewOrder(ctx context.Context, or *OrderNewRequest) (*OrderNewResp, error) {
	fmt.Println(or.OrderMain)
	fmt.Println("-------------")
	fmt.Println(or.OrderMain.OrderDetails[0])
	fmt.Println(or.OrderMain.OrderDetails[1])
	return &OrderNewResp{
		Status:  "ok",
		Message: "success",
	}, nil
}
