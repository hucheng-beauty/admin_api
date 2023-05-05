package account

import (
	"admin_api/global"
	"admin_api/internal/data"
	"admin_api/internal/pkg/password"
	"admin_api/internal/request"
	"admin_api/internal/response"
	"admin_api/internal/service/store/mysql/account"
	"admin_api/middlewares"
	"errors"
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Account struct{}

func (u *Account) Signup(c *gin.Context) {

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

func (u *Account) PasswordLogin(c *gin.Context) {
	var req *request.PasswordLogin

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, response.Error(1, err.Error()))
		return
	}

	us := account.NewUserService(data.NewUserRepo(global.DB))
	//通过username获取用户信息
	uu, err := us.GetUserByUserName(&request.GetUserByUserName{UserName: req.UserName})
	if err != nil {
		c.JSON(http.StatusOK, response.Error(1, err.Error()))
		return
	}

	//检查密码
	if ok := us.CheckPassword(&request.CheckPassword{
		Password:          req.Password,
		EncryptedPassword: uu.Password,
	}); !ok {
		c.JSON(http.StatusOK, response.Error(1, errors.New("密码错误").Error()))
		return
	}

	//添加token信息，方便下次登录直接适用token验证
	j := middlewares.NewJWT()
	token, err := j.CreateToken(middlewares.Claims{
		UserId: uu.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 60*60,
			Issuer:    "admin_api",
			NotBefore: time.Now().Unix(),
		},
	})
	if err != nil {
		c.JSON(http.StatusOK, response.Error(1, errors.New("token 生成失败").Error()))
		return
	}
	c.JSON(http.StatusOK, response.NewResponse(gin.H{"token": token}))

}

func (u *Account) GetUserDetail(c *gin.Context) {

	req := &request.GetUserDetail{}
	us := account.NewUserService(data.NewUserRepo(global.DB))
	//获取context中key是否有user_id
	userId, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusOK, response.Error(1, errors.New("获取用户Id失败").Error()))
		return
	}
	//断言用户id类型
	if uid, isString := userId.(string); !isString {
		c.JSON(http.StatusOK, response.Error(1, errors.New("获取用户Id失败").Error()))
		return
	} else {
		//绑定用户id
		req.Id = uid
	}

	//获取用户细节
	resp, err := us.GetUserDetail(req)
	if err != nil {
		//
		fmt.Println("//获取用户细节", err.Error())
		c.JSON(http.StatusOK, response.Error(1, err.Error()))
		return
	}
	//返回用户细节数据
	c.JSON(http.StatusOK, response.NewResponse(resp))
	return
}

func (Account) UpdatePassword(c *gin.Context) {
	req := &request.UpdatePassword{}
	if err := c.ShouldBind(&req); err != nil {
		fmt.Println(1)
		c.JSON(http.StatusOK, response.Error(1, err.Error()))
		return
	}

	//判断id
	userId, IsExists := c.Get("user_id")
	if !IsExists {
		fmt.Println(2)
		c.JSON(http.StatusOK, response.Error(1, errors.New("获取用户Id失败").Error()))
		return
	}
	uid, ok := userId.(string)
	if !ok {
		fmt.Println(3)
		c.JSON(http.StatusOK, response.Error(1, errors.New("内部错误").Error()))
		return
	} else {
		req.Id = uid
	}

	us := account.NewUserService(data.NewUserRepo(global.DB))
	//获取用户细节
	uu, err := us.GetUserById(&request.GetUserById{Id: req.Id})
	if err != nil {
		fmt.Println(4)
		c.JSON(http.StatusOK, response.Error(1, errors.New("内部错误").Error()))
		return
	}

	//更新密码(加密)
	req.Password = password.Generate(req.Password)

	_, err = us.UpdatePassword(req)
	if err != nil {
		fmt.Println(5)
		c.JSON(http.StatusOK, response.Error(1, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewResponse(gin.H{"id": uu.Id}))
	return
}
