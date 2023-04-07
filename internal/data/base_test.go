package data

import (
	"testing"

	"admin_api/global"
	_ "admin_api/initialize"

	"go.uber.org/zap"
)

func TestMainer(t *testing.T) {
	// simple using
	if err := global.DB.Scopes(Pagination(0, 20)).First(nil); err != nil {
		zap.S().Error()
	}
}
