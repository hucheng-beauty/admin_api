package account

//import (
//	"admin_api/internal/model"
//	"admin_api/internal/request"
//	"admin_api/pkg/uuid"
//	"github.com/go-playground/assert/v2"
//	"testing"
//)
//
//var testWalletService *Wallet
//
//func TestWallet_FrozenAmount(t *testing.T) {
//	uid := uuid.New()
//	_, err := testWalletService.CreateWallet(&model.Wallet{UserId: uid})
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	err = testWalletService.FrozenAmount(uid, 100)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	am, err := testWalletService.GetUserAmount(&request.GetUserAmount{UserId: uid})
//	if err != nil {
//		t.Error(err)
//	}
//	assert.Equal(t, am.Amount, -100)
//	assert.Equal(t, am.UsableBalance, -100)
//	assert.Equal(t, am.FrozenAmount, 100)
//
//}
