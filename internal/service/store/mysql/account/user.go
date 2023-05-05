package account

import (
	"admin_api/internal/model"
	"admin_api/internal/pkg/password"
	"admin_api/internal/request"
	"admin_api/internal/response"
	"admin_api/pkg/uuid"
	"errors"
	"go.uber.org/zap"
	"time"
)

type Repo interface {
	Save(*model.User) (*model.User, error)
	IsExist(id string, username string) bool
	Detail(user *model.User) (*model.User, error)
}

type User struct {
	repo Repo
}

func NewUserService(repo Repo) *User {
	return &User{repo: repo}
}

func (u *User) CheckPassword(req *request.CheckPassword) bool {
	return password.Verify(req.EncryptedPassword, req.Password)
}

func (u *User) IsExist(id string, username string) bool {
	return u.repo.IsExist(id, username)
}

func (u *User) CreateUser(req *request.CreateUser) (*response.CreateUser, error) {
	//绑定数据,用repo接口保存value,将前端传入的value保存到数据库
	user, err := u.repo.Save(&model.User{
		BaseModel: model.BaseModel{
			Id:        uuid.New(),
			CreatedAt: time.Now().Local(),
			UpdatedAt: time.Now().Local(),
		},
		UserName:      req.Username,
		CompanyName:   req.CompanyName,
		Password:      req.Password,
		ContactName:   req.ContactName,
		ContactMobile: req.ContactMobile,
		License:       model.File{Type: req.License.Type, URL: req.License.URL},
	})
	if err != nil {
		zap.S().Errorf("%+#v\n", err)
		return nil, errors.New("新建用户失败")
	}
	//返回创建成功的id
	return &response.CreateUser{Id: user.Id}, nil
}

func (u *User) GetUserByUserName(req *request.GetUserByUserName) (*response.GetUserByUserName, error) {
	user, err := u.repo.Detail(&model.User{
		UserName: req.UserName,
	})
	if err != nil {
		zap.S().Errorf("%+#v\n", err)
		return nil, errors.New("获取用户信息失败")
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}
	return &response.GetUserByUserName{
		Id:            user.Id,
		UserName:      user.UserName,
		Password:      user.Password,
		CompanyName:   user.CompanyName,
		ContactName:   user.ContactName,
		ContactMobile: user.ContactMobile,
		License:       response.File{Type: user.License.Type, URL: user.License.URL},
		Industry:      "user.Industry",
		Subject:       "user.Subject",
		Captcha:       user.Captcha,
		CreatedAt:     user.CreatedAt.Local().Format("2006-01-02 15:04:05"),
	}, nil
}

func (u *User) GetUserDetail(req *request.GetUserDetail) (*response.GetUserDetail, error) {
	//获取用户细节
	user, err := u.GetUserById(&request.GetUserById{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &response.GetUserDetail{
		Id:            user.Id,
		UserName:      user.UserName,
		CompanyName:   user.CompanyName,
		ContactName:   user.ContactName,
		ContactMobile: user.ContactMobile,
		License:       user.License,
		Industry:      "user.Industry",
		Subject:       "user.Subject",
		CreatedAt:     user.CreatedAt,
	}, nil
}

func (u *User) GetUserById(req *request.GetUserById) (*response.GetUserById, error) {
	user, err := u.repo.Detail(&model.User{
		BaseModel: model.BaseModel{Id: req.Id},
	})
	if err != nil {
		zap.S().Errorf("%+#v\n", err)
		return nil, errors.New("获取用户信息失败")
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}
	return &response.GetUserById{
		Id:            user.Id,
		UserName:      user.UserName,
		Password:      user.Password,
		CompanyName:   user.CompanyName,
		ContactName:   user.ContactName,
		ContactMobile: user.ContactMobile,
		License:       response.File{Type: user.License.Type, URL: user.License.URL},
		Industry:      "user.Industry",
		Subject:       "user.Subject",
		Captcha:       user.Captcha,
		CreatedAt:     user.CreatedAt.Local().Format("2006-01-02 15:04:05"),
	}, nil
}
