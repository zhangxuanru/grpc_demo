/*
@Time : 2020/5/20 16:44
@Author : zxr
@File : hystrix_test
@Software: GoLand
*/
package test

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"
)

//熔断器 练习
func TestHy(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for {
		hy := Hy()
		fmt.Println(hy)
	}
}

func Hy() string {
	location, _ := time.LoadLocation("Asia/Shanghai")
	i := rand.Intn(10)
	if i < 6 {
		time.Sleep(3 * time.Second)
	}
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	return time.Now().In(location).Format("2006-01-02 15:04:05")
}
