package request

import "fmt"

// 营销活动请求参数
type Query struct {
	Offset int    `form:"offset"`
	Limit  int    `form:"limit"`
	Filter string `form:"filter"`
}

func (q *Query) Validate() error {
	if q.Limit <= 0 || q.Offset < 0 {
		return fmt.Errorf("the Limit and Offset should not lte 0")
	}
	q.Offset = q.Offset * q.Limit
	return nil
}
