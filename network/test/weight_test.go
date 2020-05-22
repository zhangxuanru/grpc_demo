/*
@Time : 2020/5/22 17:43
@Author : zxr
@File : weight_test
@Software: GoLand
*/
package test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

type Weight struct {
	content string
	weight  int
}

type WeightList struct {
	Weight []Weight
}

//按权重输出内容
/**
根据权重值 生成类开这样的 slice,[0 0 0 0 0 1 1 2]，存储的是数组中的索引值
根据随机数生成的数在 slice中取，取出的值是对应的索引，然后取索引对应的值
**/
func TestWeight(t *testing.T) {
	list := WeightList{
		Weight: []Weight{
			{"hello1", 5},
			{"hello2", 2},
			{"hello3", 1},
		},
	}
	numList := make([]int, 0)
	randNum := 0
	for k, v := range list.Weight {
		for i := 0; v.weight > 0 && i < v.weight; i++ {
			numList = append(numList, k)
		}
		randNum += v.weight
	}
	router := mux.NewRouter()
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		rand.Seed(time.Now().Unix())
		n := rand.Intn(randNum)
		weight := list.Weight[numList[n]]
		randStr := strconv.Itoa(n)
		listStr := fmt.Sprint(numList)
		str := fmt.Sprintf("%s\n%s\n%s\n", weight.content, randStr, listStr)
		_, _ = writer.Write([]byte(str))
	})
	http.ListenAndServe(":8082", router)
}

//按权重输出内容
/**
根据权重值 生成类开这样的 slice,[4,6,7]，存储的是数组中的权重累加值
根据随机数生成的数在 slice中取，如果随机数小于slice中权重累加值，则输出权重对应的索引的内容
**/
func TestWeight2(t *testing.T) {
	list := WeightList{
		Weight: []Weight{
			{"hello1", 4},
			{"hello2", 2},
			{"hello3", 1},
		},
	}
	numList := make([]int, len(list.Weight))
	randNum := 0
	for k, v := range list.Weight {
		randNum += v.weight
		numList[k] = randNum
	}
	fmt.Println(numList)
	router := mux.NewRouter()
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		rand.Seed(time.Now().Unix())
		n := rand.Intn(randNum)
		var content string
		for k, v := range numList {
			if n < v {
				content = list.Weight[k].content
				break
			}
		}
		bytes, _ := json.Marshal(numList)
		str := fmt.Sprintf("%s\n%s\n%s\n", content, string(bytes), strconv.Itoa(n))
		writer.Write([]byte(str))
	})
	http.ListenAndServe(":8082", router)
}

func TestA(t *testing.T) {
	arr := []int{3, 5, 7}
	sprint := fmt.Sprint(arr)
	fields := strings.Fields(sprint)
	fmt.Println(strings.Join(fields, ","))

	fmt.Println(sprint)

	//strings.Replace(strings.Trim(fmt.Sprint(arr), "[]"), " ", ",", -1)

	bytes, _ := json.Marshal(arr)
	fmt.Println(string(bytes))
}
