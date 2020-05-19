/*
@Time : 2020/5/19 15:24
@Author : zxr
@File : UserTransport
@Software: GoLand
*/
package Transport

import (
	"bili/gokit/Endpoint"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DecodeUserRequest(c context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	if uid, ok := vars["uid"]; ok {
		uid, _ := strconv.Atoi(uid)
		return Endpoint.UserRequest{
			Uid: uid,
		}, nil
	}
	return nil, errors.New("params error")
}

func EncodeUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
