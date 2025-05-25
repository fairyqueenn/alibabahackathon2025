package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ZihxS/be-alibabacloud-genai-2025/constant"
	"github.com/ZihxS/be-alibabacloud-genai-2025/helper"
	"github.com/ZihxS/be-alibabacloud-genai-2025/product"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type productHandler struct {
	redisClient    *redis.Client
	productService product.Service
}

func NewProductHandler(redisClient *redis.Client, productService product.Service) *productHandler {
	return &productHandler{
		redisClient:    redisClient,
		productService: productService,
	}
}

func (handler *productHandler) GetAll(ctx *gin.Context) {
	// client, err := oss.New(constant.OSS_ENDPOINT, constant.OSS_ACCESS_KEY, constant.OSS_SECRET_KEY)
	// if err != nil {
	// 	ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
	// 	return
	// }

	// bucket, err := client.Bucket(constant.OSS_BUCKET_NAME)
	// if err != nil {
	// 	ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
	// 	return
	// }

	// _ = bucket

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"message": "success",
	// })

	products, err := handler.productService.GetAll(ctx.Request.Context())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response := helper.APIResponse(http.StatusOK, []product.Product{})
			ctx.JSON(http.StatusOK, response)
			return
		}
		response := helper.APIResponseError(http.StatusInternalServerError, err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	if products == nil {
		products = make([]*product.Product, 0)
	} else {
		client, err := oss.New(constant.OSS_ENDPOINT, constant.OSS_ACCESS_KEY, constant.OSS_SECRET_KEY)
		if err != nil {
			response := helper.APIResponseError(http.StatusInternalServerError, err.Error())
			ctx.JSON(http.StatusInternalServerError, response)
			return
		}

		bucket, err := client.Bucket(constant.OSS_BUCKET_NAME)
		if err != nil {
			response := helper.APIResponseError(http.StatusInternalServerError, err.Error())
			ctx.JSON(http.StatusInternalServerError, response)
			return
		}

		for i := range products {
			signedURLFromRedis, errRedis := getFromRedis(ctx.Request.Context(), products[i].Image, handler.redisClient)
			if errRedis == nil {
				if signedURLFromRedis != "" {
					products[i].Image = signedURLFromRedis
					continue
				}
			}
			expires := int64(60 * 5)
			signedURL, err := bucket.SignURL(products[i].Image, oss.HTTPGet, expires)
			if err != nil {
				response := helper.APIResponseError(http.StatusInternalServerError, err.Error())
				ctx.JSON(http.StatusInternalServerError, response)
				return
			}
			imageKey := products[i].Image
			products[i].Image = signedURL
			_ = setToRedis(ctx.Request.Context(), imageKey, products[i].Image, (time.Second*60*5)-10, handler.redisClient)
		}
	}

	response := helper.APIResponse(http.StatusOK, products)
	ctx.JSON(http.StatusOK, response)
}

func (handler *productHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		response := helper.APIResponseError(http.StatusBadRequest, err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := handler.productService.GetByID(ctx.Request.Context(), intID)
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

	signedURLFromRedis, errRedis := getFromRedis(ctx.Request.Context(), product.Image, handler.redisClient)
	if errRedis == nil {
		if signedURLFromRedis != "" {
			product.Image = signedURLFromRedis
		}
	}

	if signedURLFromRedis == "" {
		client, err := oss.New(constant.OSS_ENDPOINT, constant.OSS_ACCESS_KEY, constant.OSS_SECRET_KEY)
		if err != nil {
			response := helper.APIResponseError(http.StatusInternalServerError, err.Error())
			ctx.JSON(http.StatusInternalServerError, response)
			return
		}

		bucket, err := client.Bucket(constant.OSS_BUCKET_NAME)
		if err != nil {
			response := helper.APIResponseError(http.StatusInternalServerError, err.Error())
			ctx.JSON(http.StatusInternalServerError, response)
			return
		}

		expires := int64(60 * 5)
		signedURL, err := bucket.SignURL(product.Image, oss.HTTPGet, expires)
		if err != nil {
			response := helper.APIResponseError(http.StatusInternalServerError, err.Error())
			ctx.JSON(http.StatusInternalServerError, response)
			return
		}
		imageKey := product.Image
		product.Image = signedURL
		_ = setToRedis(ctx.Request.Context(), imageKey, product.Image, (time.Second*60*5)-10, handler.redisClient)
	}

	response := helper.APIResponse(http.StatusOK, product)
	ctx.JSON(http.StatusOK, response)
}

func (handler *productHandler) Add(ctx *gin.Context) {
	var req product.AddProductRequest

	err := ctx.ShouldBind(&req)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponseError(http.StatusUnprocessableEntity, errors[0])
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	fmt.Printf("REQ %+v\n", req)

	client, err := oss.New(constant.OSS_ENDPOINT, constant.OSS_ACCESS_KEY, constant.OSS_SECRET_KEY)
	if err != nil {
		response := helper.APIResponseError(http.StatusInternalServerError, err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	bucket, err := client.Bucket(constant.OSS_BUCKET_NAME)
	if err != nil {
		response := helper.APIResponseError(http.StatusInternalServerError, err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		response := helper.APIResponseError(http.StatusBadRequest, err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	defer file.Close()

	objectName := fmt.Sprintf("uploads/%s-%s", uuid.New().String(), header.Filename)
	err = bucket.PutObject(objectName, file)
	if err != nil {
		response := helper.APIResponseError(http.StatusInternalServerError, err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	req.Image = &product.ProductUploadedFile{
		FileName: objectName,
		MIMEType: header.Header.Get("Content-Type"),
		Size:     header.Size,
		Content:  nil,
	}

	product, err := handler.productService.AddProduct(ctx.Request.Context(), req)
	if err != nil {
		response := helper.APIResponseError(http.StatusInternalServerError, err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	ctx.JSON(http.StatusCreated, helper.APIResponse(http.StatusCreated, gin.H{
		"product_id": product.ID,
	}))
}

func (handler *productHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		response := helper.APIResponseError(http.StatusBadRequest, err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var req product.UpdateProductRequest
	if err := ctx.ShouldBind(&req); err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponseError(http.StatusUnprocessableEntity, errors[0])
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = handler.productService.UpdateProduct(ctx.Request.Context(), intID, req)
	if err != nil {
		response := helper.APIResponseError(http.StatusInternalServerError, err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
}
