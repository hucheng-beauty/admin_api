package marketing

import (
	"admin_api/global"
	"admin_api/internal/data"
	"admin_api/internal/service/store/mysql/account"
)

func NewTradeService() *account.TradeService {
	return account.NewTradeService(data.NewTradeRepo(global.DB))
}
