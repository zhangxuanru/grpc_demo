/*
@Time : 2020/5/20 15:08
@Author : zxr
@File : main_rate
@Software: GoLand
*/
package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

//令牌桶
func main() {
	r := rate.NewLimiter(1, 5)
	ctx := context.Background()
	for {
		_ = r.WaitN(ctx, 2) //阻塞等待
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(1 * time.Second)
	}
}
