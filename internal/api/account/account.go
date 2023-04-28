package account

import (
	"admin_api/global"
	"admin_api/internal/data"
	"admin_api/internal/pkg/password"
	"admin_api/internal/request"
	"admin_api/internal/response"
	"admin_api/internal/service/store/mysql/account"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Account struct{}

func (Account) Signup(c *gin.Context) {

	var req *request.CreateUser

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, response.Error(1, err.Error()))
		//
		fmt.Println("1", err.Error())
		return
	}

	us := account.NewUserService(data.NewUserRepo(global.DB))
	//判断用户是否存在,调用方法 方法调用接口
	if ok := us.IsExist("", req.Username); ok {
		c.JSON(http.StatusOK, response.Error(1, errors.New("用户已存在").Error()))
		return
	}

	//加密密码
	req.Password = password.Generate(req.Password)
	//创建用户
	resp, err := us.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusOK, response.Error(1, err.Error()))
		//
		fmt.Println("2", err.Error())
		return
	}

	c.JSON(http.StatusOK, response.NewResponse(resp))
	return
}

func (Account) PasswordLogin(c *gin.Context) {

}

func (Account) GetUserDetail(c *gin.Context) {

}

func (Account) UpdatePassword(c *gin.Context) {

}
