package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

const TotalCountHeader = `X-Total-Count`

func TotalCount(ctx *gin.Context, totalCount int) {
	ctx.Writer.Header().Set(TotalCountHeader, fmt.Sprintf("%d", totalCount))
}
