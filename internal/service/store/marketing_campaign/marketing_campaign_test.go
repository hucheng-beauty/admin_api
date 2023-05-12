package marketing_campaign

import (
	"admin_api/internal/model"
	"admin_api/internal/request"
	"admin_api/pkg/uuid"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

var testMarCampaignService *MarCampaignService

//func TestMain(t *testing.M) {
//	dir, _ := os.Getwd()
//	fmt.Println(dir)
//	test.Init("../../../../test/etc/config_test.yaml")
//	err := test.DB.AutoMigrate(
//		model.CouponBatch{},
//		model.CouponTemplate{},
//		model.MarketingCampaign{},
//		model.CouponLog{},
//		model.MarketingCampaignLog{},
//		model.Trade{},
//	)
//	if err != nil {
//		panic(err)
//	}
//	testMarCampaignService = NewMarCampaignService(
//		data.NewMarketingCampaignRepo(test.DB),
//		data.NewCouponBatchRepo(test.DB),
//		data.NewCouponTemplateRepo(test.DB),
//		data.NewCouponLogRepo(test.DB),
//		data.NewMarketingCampaignLogRepo(test.DB),
//	)
//
//	os.Exit(t.Run())
//
//}
//
//func TestRun(t *testing.T) {
//
//}

func createTestMarCampaign(t *testing.T) (*model.MarketingCampaign, []*model.CouponBatch) {
	var tmp = &model.CouponTemplate{}
	tmp.StockName = uuid.New()    //批次名称
	tmp.MerchantName = uuid.New() // 服务商名称
	tmp.TransactionMinimum = 310  // 使用限制(门槛)，单位分,必须大于优惠金额
	tmp.CouponAmount = 300        // 面额,单位分，单位分,整数;1元<面额<=500元
	tmp.Platform = 1              // 所属平台，枚举(1-支付宝、2-微信)
	tmp.Type = 1                  // 枚举(1-满减)
	// 创建模板
	template, err := testMarCampaignService.CreateCouponTemplate(tmp)
	if err != nil {
		t.Error(err)
		return nil, nil
	}

	var req = &request.CreateMarketingCampaignRequest{}

	req.CampaignName = uuid.New()

	req.AvailableBeginTime = time.Now()
	req.AvailableEndTime = time.Now()

	var couponBatch = request.CreateCouponBatchReq{}
	couponBatch.TemplateID = template.Id
	couponBatch.Comment = uuid.New()
	couponBatch.Bin = []string{"123", "1233"}

	req.CouponBatches = append(req.CouponBatches, couponBatch)

	// 创建活动和券批次
	mr, cbs, err := testMarCampaignService.CreateMarCampaignAndCouponBatch(req, uuid.New())
	if err != nil {
		t.Error(err)
		return nil, nil
	}

	assert.Equal(t, mr.CampaignName, req.CampaignName)
	assert.Equal(t, mr.AvailableBeginTime, req.AvailableBeginTime)
	assert.Equal(t, mr.AvailableEndTime, req.AvailableEndTime)

	// 查券批次是不是用了模板信息
	cbsInDB, err := testMarCampaignService.CouponBatchByMarId(mr.Id)
	if err != nil {
		t.Error(err)
		return nil, nil
	}
	assert.Equal(t, len(cbsInDB), 1)
	assert.Equal(t, cbsInDB[0].StockName, tmp.StockName)
	assert.Equal(t, cbsInDB[0].MerchantName, tmp.MerchantName)
	assert.Equal(t, cbsInDB[0].Comment, couponBatch.Comment)

	return mr, cbs
}

// 创建营销活动
func TestMarCampaignService_CreateMarCampaign(t *testing.T) {
	createTestMarCampaign(t)
}

func TestMarCampaignService_UpdateMarCampaignStateById(t *testing.T) {
	cm, _ := createTestMarCampaign(t)
	err := testMarCampaignService.UpdateMarCampaignStateById(cm.Id, -1)
	if err != nil {
		t.Error(err)
		return
	}
	editData, err := testMarCampaignService.MarCampaignById(cm.Id)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, editData.State, -1)
}

func TestMarCampaignService_UpdateMarCampaignCouponSurplusNumber(t *testing.T) {
	cm, _ := createTestMarCampaign(t)

	success := rand.Intn(1000) + 1

	err := testMarCampaignService.UpdateMarCampaignCouponSurplusNumber(cm.Id, success)
	if err != nil {
		t.Error(err)
		return
	}
	nCm, err := testMarCampaignService.MarCampaignById(cm.Id)

	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, nCm.CouponSurplusNumber, cm.CouponSurplusNumber-int64(success))

}

// 卷模板
func TestCouponTemplate(t *testing.T) {
	var tmp = &model.CouponTemplate{}
	tmp.StockName = "测试批次名称"     //批次名称
	tmp.MerchantName = "测试服务商名称" // 服务商名称
	tmp.TransactionMinimum = 310 // 使用限制(门槛)，单位分,必须大于优惠金额
	tmp.CouponAmount = 300       // 面额,单位分，单位分,整数;1元<面额<=500元
	tmp.Platform = 1             // 所属平台，枚举(1-支付宝、2-微信)
	tmp.Type = 1                 // 枚举(1-满减)
	// 创建模板
	ret, err := testMarCampaignService.CreateCouponTemplate(tmp)
	fmt.Println(ret, err)
}
