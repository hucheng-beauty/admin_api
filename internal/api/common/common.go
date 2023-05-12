package common

import (
	"admin_api/internal/request"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func GetUserId(c *gin.Context) (string, error) {
	userId, ok := c.Get("user_id")
	if !ok {
		return "", errors.New("用户Id不存在")
	}

	uid, yeah := userId.(string)
	if !yeah {
		return "", errors.New("内部错误")
	}
	return uid, nil
}

func OffsetAndLimitHandle(c *gin.Context) request.Pagination {
	var (
		err error
		p   request.Pagination
	)

	offset := c.Query("offset")
	if offset != "" {
		p.Offset, err = strconv.Atoi(c.Query("offset"))
		if err != nil {
			zap.S().Error(zap.String("offset error:", err.Error()))
			p.Offset = 0
		}
	}

	limit := c.Query("limit")
	if limit != "" {
		p.Limit, err = strconv.Atoi(c.Query("limit"))
		if err != nil {
			zap.S().Error(zap.String("limit error:", err.Error()))
			p.Limit = 20
		}
	}

	if p.Offset <= 0 {
		p.Offset = 0
	}
	if p.Limit <= 0 {
		p.Limit = 20
	}
	//if p.Limit >= 50 { // TODO 确认是否需要限制
	//	p.Limit = 50
	//}
	if p.Offset != 0 {
		p.Offset = p.Offset * p.Limit
	}
	return p
}
