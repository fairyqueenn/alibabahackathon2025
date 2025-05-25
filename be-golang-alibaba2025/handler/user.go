package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/ZihxS/be-alibabacloud-genai-2025/helper"
	"github.com/ZihxS/be-alibabacloud-genai-2025/user"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type userHandler struct {
	redisClient *redis.Client
	userService user.Service
}

func NewUserHandler(redisClient *redis.Client, userService user.Service) *userHandler {
	return &userHandler{
		redisClient: redisClient,
		userService: userService,
	}
}

func (handler *userHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		response := helper.APIResponseError(http.StatusBadRequest, err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := handler.userService.GetUserByID(ctx.Request.Context(), intID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response := helper.APIResponseError(http.StatusNotFound, err.Error())
			ctx.JSON(http.StatusNotFound, response)
			return
		}
		response := helper.APIResponseError(http.StatusInternalServerError, err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.APIResponse(http.StatusOK, user)
	ctx.JSON(http.StatusOK, response)
}
