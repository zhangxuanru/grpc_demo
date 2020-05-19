/*
@Time : 2020/5/18 19:13
@Author : zxr
@File : UserService
@Software: GoLand
*/
package Service

type IUserService interface {
	GetName(uid int) string
}

type UserService struct {
}

func (u *UserService) GetName(uid int) string {
	return "yes"
}
