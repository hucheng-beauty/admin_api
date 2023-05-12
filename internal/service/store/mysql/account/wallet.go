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

type WalletRepo interface {
	Save(wallet *model.Wallet) (*model.Wallet, error)
	Detail(wallet *model.Wallet) (*model.Wallet, error)
	CheckId(id string) (bool, error)
	Update(id string) (*model.Wallet, error)
}

type Wallet struct {
	repo WalletRepo
}

func NewWalletService(repo WalletRepo) *Wallet {
	return &Wallet{repo: repo}
}

func (w *Wallet) GetUserAmount(req *request.GetUserAmount) (*response.GetUserAmount, error) {
	detail, err := w.repo.Detail(&model.Wallet{UserId: req.UserId})

	if err != nil {
		zap.S().Errorf("%+#v\n", err)
		return nil, errors.New("查询用户钱包失败")
	}

	if detail == nil {
		return nil, errors.New("该用户钱包不存在")
	}
	return &response.GetUserAmount{
		UserId:        detail.UserId,
		Amount:        detail.Amount,
		UsableBalance: detail.UsableBalance,
		FrozenAmount:  detail.FrozenAmount,
	}, nil
}

func (w *Wallet) CreateWallet(wallet *model.Wallet) (*response.GetAmountId, error) {
	wallet.BaseModel = model.BaseModel{
		Id:        uuid.New(),
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
	}
	//wallet.UserId = id
	ww, err := w.repo.Save(wallet)
	if err != nil {
		return nil, errors.New("用户钱包创建失败")
	}
	return &response.GetAmountId{Userid: ww.Id}, err
}

func (w *Wallet) CheckWalletUserId(id string) (bool, error) {
	return w.repo.CheckId(id)
}

func (w *Wallet) UpdateWallet(id string) (*response.GetAmountId, error) {
	ww, err := w.repo.Update(id)
	return &response.GetAmountId{Userid: ww.Id}, err
}
