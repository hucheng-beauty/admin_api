package publish

import (
	"admin_api/internal/request"
	"context"
	"fmt"
	"github.com/maocatooo/thin/broker"
)

var (
	marketingCampaignSendRecordEvent broker.Event
)

func WithMarketingCampaignSendRecord(ctx context.Context, data *request.UpdateSendRecord) {
	err := marketingCampaignSendRecordEvent.Publish(ctx, data)
	if err != nil {
		fmt.Println(err)
	}
}
