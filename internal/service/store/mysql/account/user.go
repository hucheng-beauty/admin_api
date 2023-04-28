package account

import (
	"admin_api/internal/model"
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
}

type User struct {
	repo Repo
}

func NewUserService(repo Repo) *User {
	return &User{repo: repo}
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
