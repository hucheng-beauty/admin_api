package publish

import (
	"admin_api/internal/model"
	"context"
	"fmt"
	"github.com/maocatooo/thin/broker"
)

var (
	marketingCampaignCreateEvent broker.Event
)

// 活动创建
func WithMarketingCampaignCreate(ctx context.Context, campaign *model.MarketingCampaign) {
	if campaign == nil {
		return
	}
	err := marketingCampaignCreateEvent.Publish(ctx, campaign)
	if err != nil {
		fmt.Println(err)
	}
}
