package financial

import (
	"admin_api/global"
	"admin_api/internal/api/common"
	"admin_api/internal/data"
	"admin_api/internal/model"
	"admin_api/internal/request"
	"admin_api/internal/response"
	"admin_api/internal/service/store/mysql/account"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Financial struct {
}

// GetUserAmount 获取账户可用金额和冻结金额
func (f Financial) GetUserAmount(c *gin.Context) {

	uid, err := common.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusOK, response.Error(1, err.Error()))
		return
	}

	//初始化数据库
	ws := account.NewWalletService(data.NewWalletService(global.DB))
	//获取用户钱包Wallet struct
	resp, err := ws.GetUserAmount(&request.GetUserAmount{UserId: uid})
	if err != nil {
		c.JSON(http.StatusOK, response.Error(1, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.NewResponse(resp))
	return
}

// 创建用户钱包
func (f Financial) CreateUserAmount(c *gin.Context) {
	id, err := common.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusOK, response.Error(1, err.Error()))
		return
	}

	//初始化数据库
	ws := account.NewWalletService(data.NewWalletService(global.DB))

	todo, err := ws.CheckWalletUserId(id)
	if err != nil {
		c.JSON(http.StatusOK, response.Error(1, err.Error()))
		return
	}
	if todo == true {
		resp, err := ws.UpdateWallet(id)
		if err != nil {
			c.JSON(http.StatusOK, response.Error(1, err.Error()))
			return
		}
		c.JSON(http.StatusOK, response.NewResponse(resp))
	} else {
		resp, err := ws.CreateWallet(&model.Wallet{UserId: id, Amount: 1000000000})
		if err != nil {
			c.JSON(http.StatusOK, response.Error(1, err.Error()))
			return
		}
		c.JSON(http.StatusOK, response.NewResponse(resp))
	}

}

//// DescribeUserTrade 获取账户下的交易流水
//func (f Financial) DescribeUserTrade(c *gin.Context) {
//	uid, err := common.GetUserId(c)
//	if err != nil {
//		c.JSON(http.StatusOK, response.Error(1, err.Error()))
//		return
//	}
//
//	//初始化数据库
//	ts := account.NewTradeService(data.NewTradeRepo(global.DB))
//	resp, total, err := ts.DescribeUserTrade(&request.DescribeUserTrade{
//		Pagination:   common.OffsetAndLimitHandle(c),
//		FilterString: request.FilterString{Filter: c.Query("filter")},
//		UserId:       uid,
//	})
//	if err != nil {
//		c.JSON(http.StatusOK, response.Error(1, err.Error()))
//		return
//	}
//	c.JSON(http.StatusOK, response.SuccessWithPagination(resp, response.NewPagination(total)))
//	return
//}
