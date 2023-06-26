package marketing

import (
	"admin_api/internal/request"
	"admin_api/internal/response"
	"github.com/gin-gonic/gin"
)

type CouponBatchApi struct{}

func (CouponBatchApi) Logs(c *gin.Context, in *request.Query, out *response.CouponLogsResponse) error {

	err := in.Validate()
	if err != nil {
		return err
	}
	out.Data, out.TotalCount, err = NewMarCampaignService().ListWitPage(c.Param("id"), in)
	if err != nil {
		return err
	}

	return nil
}
