/*
@Time : 2020/5/20 18:25
@Author : zxr
@File : jwt_test
@Software: GoLand
*/
package test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/dgrijalva/jwt-go"
)

//jwt test

type UserInfo struct {
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

//对称加密
func TestJwt1(t *testing.T) {
	sec := []byte("123abc")
	//生成token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, UserInfo{UserName: "zxr"})
	token, _ := claims.SignedString(sec)
	fmt.Println(token)

	//token比对
	parse, _ := jwt.Parse(token, func(token *jwt.Token) (i interface{}, e error) {
		return sec, nil
	})

	/*
		//解析成结构体
		user := &UserInfo{}
		withClaims, _ := jwt.ParseWithClaims(token, user, func(token *jwt.Token) (i interface{}, e error) {
			return sec, nil
		})
		if withClaims.Valid {
			fmt.Println(withClaims.Claims.(*UserInfo))
		}
	*/

	if parse.Valid {
		fmt.Println(parse.Claims)
	} else {
		fmt.Println("token error.....")
	}
}

//非对称加密
func TestJwt2(t *testing.T) {
	_ = jwt2GenPemFile(1024, "pem")
	priBytes, _ := ioutil.ReadFile("./pem/priv.pem")
	pubBytes, _ := ioutil.ReadFile("./pem/pub.pem")

	privkey, _ := jwt.ParseRSAPrivateKeyFromPEM(priBytes)

	pubKey, _ := jwt.ParseRSAPublicKeyFromPEM(pubBytes)

	claims := jwt.NewWithClaims(jwt.SigningMethodRS256, UserInfo{UserName: "zxr good good "})
	token, _ := claims.SignedString(privkey)

	fmt.Println("token:", token)

	user := &UserInfo{}
	withClaims, _ := jwt.ParseWithClaims(token, user, func(token *jwt.Token) (i interface{}, e error) {
		return pubKey, nil
	})
	if withClaims.Valid {
		fmt.Println(withClaims.Claims)
	} else {
		fmt.Println("token error......")
	}
}

//生成公钥私钥文件
func jwt2GenPemFile(bits int, path string) error {
	privateKey, _ := rsa.GenerateKey(rand.Reader, bits)
	key := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: key,
	}
	err := ioutil.WriteFile(path+"/priv.pem", pem.EncodeToMemory(block), os.ModePerm)
	publicKey := &privateKey.PublicKey
	bytes, _ := x509.MarshalPKIXPublicKey(publicKey)
	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: bytes,
	}
	_ = ioutil.WriteFile(path+"/pub.pem", pem.EncodeToMemory(publicBlock), os.ModePerm)
	return err
}
