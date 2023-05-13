package wechat

import (
	"context"
	"crypto/rsa"
	"fmt"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"time"
)

type Info struct {
	mchID                      string
	mchCertificateSerialNumber string
	mchAPIv3Key                string
	appid                      string
	appSecret                  string
	privateKeyPath             string
	cache                      cache.Cache
}

// 构建基础信息
func NewDefaultInfo() *Info {
	var (
		mchID                      string = "xxxxxxxxxx"                               // 商户号
		mchCertificateSerialNumber string = "xxxxxxxxxx111111111111111111111111111111" // 商户证书序列号
		mchAPIv3Key                string = "xxxxxxxxxx1111111111111111111111"         // 商户 APIv3 密钥
		appid                      string = "xxxxxxxxxx11111111"                       // 应用 Id
		appSecret                  string = "xxxxxxxxxx1111111111111111111111"         // app 密钥
		filepath                   string = `./xxx/apiclient_key.pem`                  // 商户 API 私钥
	)
	return NewInfo(mchID, mchCertificateSerialNumber, mchAPIv3Key, appid, appSecret, filepath)
}

func NewInfo(mchID, mchCertificateSerialNumber, mchAPIv3Key, appid, appSecret, privateKeyPath string) *Info {
	return &Info{mchID: mchID, mchCertificateSerialNumber: mchCertificateSerialNumber,
		mchAPIv3Key: mchAPIv3Key, appid: appid, appSecret: appSecret,
		cache: cache.NewMemory(), privateKeyPath: privateKeyPath}
}

// Send 发放券批次:https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_2.shtml
func (info *Info) Send(StockId, OpenId string) (*SendResp, error) {
	client, err := info.GetClient()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.mch.weixin.qq.com/v3/marketing/favor/users/%s/coupons", OpenId)
	resp, err := client.Post(context.Background(), url, &SendReq{
		StockId:           StockId,
		OutRequestNo:      fmt.Sprintf("%d", time.Now().Nanosecond()),
		AppId:             info.appid,
		StockCreatorMchid: info.mchID,
	})
	if err != nil {
		return nil, err
	}

	sendResp := &SendResp{}
	err = core.UnMarshalResponse(resp.Response, sendResp)
	return sendResp, err
}

func (info *Info) GetClient() (client *core.Client, err error) {
	var getApiClientPrivateKey = func() *rsa.PrivateKey {
		rt, _ := utils.LoadPrivateKeyWithPath(info.privateKeyPath)
		return rt
	}
	return core.NewClient(context.Background(), []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(info.mchID, info.mchCertificateSerialNumber,
			getApiClientPrivateKey(), info.mchAPIv3Key)}...)
}

// GetCouponDetail 查询代金券详情:https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_6.shtml
func (info *Info) GetCouponDetail(CouponId, OpenId string) (*GetCouponDetailResp, error) {
	client, err := info.GetClient()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.mch.weixin.qq.com/v3/marketing/favor/users/%s/coupons/%s?appid=%s",
		OpenId, CouponId, info.appid)
	resp, err := client.Get(context.Background(), url)
	if err != nil {
		return nil, err
	}
	getResp := &GetCouponDetailResp{}
	err = core.UnMarshalResponse(resp.Response, getResp)
	return getResp, err
}
