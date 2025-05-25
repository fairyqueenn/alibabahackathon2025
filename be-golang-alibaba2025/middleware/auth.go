package middleware

import (
	"net/http"
	"slices"

	"github.com/ZihxS/be-alibabacloud-genai-2025/helper"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("X-API-Key")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.BasicAPIResponse(http.StatusUnauthorized))
			return
		}
		allowedKeyAPI := []string{"wx9bHXTUDo", "d3g9heg3xh", "Po3kFH06VR", "E4mavBRTFl", "DoLUoh6tBq"}
		if !slices.Contains(allowedKeyAPI, authHeader) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.BasicAPIResponse(http.StatusUnauthorized))
			return
		}
	}
}
