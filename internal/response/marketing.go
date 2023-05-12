package response

import "admin_api/internal/model"

type Id struct {
	Id string `json:"id"`
}

type MarketingCampaignListResponse struct {
	TotalCount int `json:"total_count"`

	Data []*model.MarketingCampaign `json:"data"`
}
