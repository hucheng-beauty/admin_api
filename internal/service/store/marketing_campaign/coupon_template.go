package marketing_campaign

import (
	"admin_api/internal/model"
	"fmt"
)

func (s *MarCampaignService) CreateCouponTemplate(ct *model.CouponTemplate) (*model.CouponTemplate, error) {
	ct.BM = NewBM()
	return ct, s.ctr.Create(ct)
}

// 券批次所用的模板Map {"template_id": template}
func (s *MarCampaignService) CouponBatchTemplates2Map(cbs []*model.CouponBatch) (map[string]*model.CouponTemplate, error) {
	ids := make([]string, 0)
	for _, item := range cbs {
		if item.TemplateID == `` {
			return nil, fmt.Errorf("the template_id should not nil")
		}

		ids = append(ids, item.TemplateID)
	}
	//根据ids获取卷模板
	res, err := s.ctr.FindByIds(ids)
	if err != nil {
		return nil, err
	}
	mp := map[string]*model.CouponTemplate{}
	//mp存储卷模板id和对应卷模板
	for _, item := range res {
		mp[item.Id] = item
	}
	return mp, err
}
