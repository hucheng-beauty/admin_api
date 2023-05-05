package account

type TradeRepo interface {
}

type TradeService struct {
	repo TradeRepo
}

func NewTradeService(repo TradeRepo) *TradeService {
	return &TradeService{repo: repo}
}
