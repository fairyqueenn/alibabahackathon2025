package handler

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type aiHandler struct {
	redisClient *redis.Client
}

func NewAIHandler(redisClient *redis.Client) *aiHandler {
	return &aiHandler{redisClient: redisClient}
}

func getFromRedis(ctx context.Context, key string, redisClient *redis.Client) (string, error) {
	val, err := redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return val, nil
}

func setToRedis(ctx context.Context, key string, value string, expiration time.Duration, redisClient *redis.Client) error {
	return redisClient.Set(ctx, key, value, expiration).Err()
}

// func (handler *aiHandler) Generate(ctx *gin.Context) {
// 	var req ai.RequestGenerate

// 	err := ctx.ShouldBind(&req)

// 	if err != nil {
// 		errors := helper.FormatValidationError(err)
// 		response := helper.APIResponseError(http.StatusUnprocessableEntity, errors[0])
// 		ctx.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	valueFromRedis, _ := getFromRedis(ctx.Request.Context(), "?", handler.redisClient)
// 	if valueFromRedis != "" {
// 		ctx.JSON(http.StatusOK, helper.APIResponse(http.StatusOK, map[string]any{
// 			"output": valueFromRedis,
// 		}))
// 		return
// 	}

// 	b, err := json.Marshal(req)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	summary := ""

// 	IsIncludeCutlery := gjson.Get(string(b), "IsIncludeCutlery").Bool()
// 	Items := gjson.Get(string(b), "Items")

// 	idx := 1

// 	Items.ForEach(func(key, value gjson.Result) bool {
// 		primaries := ""
// 		cutleries := ""
// 		composables := ""
// 		summary = summary + "\n" + "Wastes:"
// 		Quantity := value.Get("Quantity")
// 		Wastes := value.Get("Wastes")
// 		Wastes.ForEach(func(key, value gjson.Result) bool {
// 			switch value.Get("Type").String() {
// 			case "primary":
// 				primaries = primaries + " " + Quantity.String() + " " + value.Get("Name").String() + " (" + value.Get("Category").String() + ")" + ","
// 			case "cutlery":
// 				if IsIncludeCutlery {
// 					cutleries = cutleries + " " + Quantity.String() + " " + value.Get("Name").String() + " (" + value.Get("Category").String() + ")" + ","
// 				}
// 			case "composable":
// 				composables = composables + " " + Quantity.String() + " " + value.Get("Name").String() + " (" + value.Get("Category").String() + ")" + ","
// 			}
// 			return true
// 		})
// 		isCutleriesDoneGenerate := false
// 		if primaries != "" {
// 			primaries = strings.TrimSpace(primaries)
// 			primaries = strings.Trim(primaries, ",")
// 			summary = summary + "\nNon Composable: " + primaries
// 			if cutleries != "" {
// 				cutleries = strings.TrimSpace(cutleries)
// 				cutleries = strings.Trim(cutleries, ",")
// 				summary = summary + "," + cutleries
// 				isCutleriesDoneGenerate = true
// 			}
// 		}
// 		if cutleries != "" && !isCutleriesDoneGenerate {
// 			cutleries = strings.TrimSpace(cutleries)
// 			cutleries = strings.Trim(cutleries, ",")
// 			summary = summary + "\nNon Composable: " + cutleries
// 		}
// 		if composables != "" {
// 			composables = strings.TrimSpace(composables)
// 			composables = strings.Trim(composables, ",")
// 			summary = summary + "\nComposable: " + composables
// 		}
// 		summary = summary + "\n"
// 		idx++
// 		return true
// 	})

// 	AdditionalWastes := gjson.Get(string(b), "AdditionalWastes")
// 	if AdditionalWastes.String() != "" {
// 		AdditionalWastes.ForEach(func(key, value gjson.Result) bool {
// 			summary = summary + "\n" + "Additional Wastes: " + value.Get("Quantity").String() + " " + value.Get("Name").String() + " (" + value.Get("Category").String() + ")"
// 			return true
// 		})
// 	}

// 	summary = strings.ReplaceAll(summary, "\n\n", "\n")
// 	summary = strings.Trim(summary, "\n")

// 	preOutput := ""

// 	// pre ai
// 	{
// 		url := "https://dashscope-intl.aliyuncs.com/api/v1/apps/71a65784d55a4d3a8a963418a9d5739d/completion"

// 		requestBody := map[string]any{
// 			"input": map[string]string{
// 				"prompt": "analogikan seberat atau seluas atau setinggi apa jumlah produksi sampah per hari di bandung?, buat sesimple mungkin dan contohkan hanya satu contoh saja, message length jangan lebih dari 250, gunakan satuan ton.",
// 			},
// 			"parameters": map[string]any{},
// 			"debug":      map[string]any{},
// 		}

// 		jsonData, err := json.Marshal(requestBody)
// 		if err != nil {
// 			log.Fatalf("error marshaling JSON: %v", err)
// 		}

// 		request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
// 		if err != nil {
// 			log.Fatalf("error creating request: %v", err)
// 		}

// 		request.Header.Add("Authorization", "Bearer "+apiKey)
// 		request.Header.Add("Content-Type", "application/json")

// 		client := &http.Client{}
// 		resp, err := client.Do(request)
// 		if err != nil {
// 			log.Fatalf("error sending request: %v", err)
// 		}
// 		defer resp.Body.Close()

// 		body, err := io.ReadAll(resp.Body)
// 		if err != nil {
// 			log.Fatalf("error reading response body: %v", err)
// 		}

// 		preOutput = gjson.Get(string(body), "output.text").String()
// 	}

// 	output := ""
// 	{
// 		url := "https://dashscope-intl.aliyuncs.com/api/v1/apps/9f9687b76e76444890c02668dc864205/completion"
// 		requestBody := map[string]any{
// 			"input": map[string]string{
// 				"prompt": "karena \"" + preOutput + "\"" + ", untuk mendukung upaya mengurangi sampah, apa prakarya atau composting yang dapat saya buat dari bahan-bahan berikut yang sudah saya beli dari delivery food app: " + summary,
// 			},
// 			"parameters": map[string]any{},
// 			"debug":      map[string]any{},
// 		}

// 		jsonData, err := json.Marshal(requestBody)
// 		if err != nil {
// 			log.Fatalf("error marshaling JSON: %v", err)
// 		}

// 		request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
// 		if err != nil {
// 			log.Fatalf("error creating request: %v", err)
// 		}

// 		request.Header.Add("Authorization", "Bearer "+apiKey)
// 		request.Header.Add("Content-Type", "application/json")

// 		client := &http.Client{}
// 		resp, err := client.Do(request)
// 		if err != nil {
// 			log.Fatalf("error sending request: %v", err)
// 		}
// 		defer resp.Body.Close()

// 		body, err := io.ReadAll(resp.Body)
// 		if err != nil {
// 			log.Fatalf("error reading response body: %v", err)
// 		}
// 		output = gjson.Get(string(body), "output.text").String()
// 	}

// 	// output = preOutput + "\n\n" + output

// 	output = strings.ReplaceAll(output, "dis analogikan", "dianalogikan")
// 	_ = setToRedis(ctx.Request.Context(), "?", output, time.Second*10, handler.redisClient)

// 	ctx.JSON(http.StatusOK, helper.APIResponse(http.StatusOK, map[string]any{
// 		"preOutput": preOutput,
// 		"output":    output,
// 	}))
// }

// func (handler *aiHandler) GenerateByImage(ctx *gin.Context) {
// 	var req ai.RequestGenerateByImage

// 	err := ctx.ShouldBind(&req)

// 	if err != nil {
// 		errors := helper.FormatValidationError(err)
// 		response := helper.APIResponseError(http.StatusUnprocessableEntity, errors[0])
// 		ctx.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	valueFromRedis, _ := getFromRedis(ctx.Request.Context(), req.Image, handler.redisClient)
// 	if valueFromRedis != "" {
// 		ctx.JSON(http.StatusOK, helper.APIResponse(http.StatusOK, map[string]any{
// 			"output": valueFromRedis,
// 		}))
// 		return
// 	}

// 	url := "https://dashscope-intl.aliyuncs.com/api/v1/services/aigc/multimodal-generation/generation"

// 	requestBody := map[string]interface{}{
// 		"model": "qwen-vl-max",
// 		"input": map[string]interface{}{
// 			"messages": []map[string]interface{}{
// 				{
// 					"role": "system",
// 					"content": []map[string]string{
// 						{"text": "AI ini adalah seorang pengerajin prakarya yang sering menggunakan bahan-bahan daur ulang. AI ini juga adalah seorang ahli dalam mengolah composable dari makanan. Jika composable tidak usah didaur ulang, berikan saja saran terbaik untuk melakukan composablenya."},
// 					},
// 				},
// 				{
// 					"role": "user",
// 					"content": []map[string]string{
// 						{"image": req.Image},
// 						{"text": "Ada wastes dan compostable apa aja di gambar tersebut?, apa prakarya atau composting yang dapat saya buat dari wastes dan compostables tersebut?"},
// 					},
// 				},
// 			},
// 		},
// 		"parameters": map[string]interface{}{},
// 	}

// 	jsonData, err := json.Marshal(requestBody)
// 	if err != nil {
// 		log.Fatalf("error marshaling JSON: %v", err)
// 	}

// 	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		log.Fatalf("error creating request: %v", err)
// 	}

// 	request.Header.Add("Authorization", "Bearer "+accessToken)
// 	request.Header.Add("Content-Type", "application/json")
// 	request.Header.Add("X-DashScope-SSE", "enable")

// 	client := &http.Client{}
// 	resp, err := client.Do(request)
// 	if err != nil {
// 		log.Fatalf("error sending request: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalf("error reading response body: %v", err)
// 	}

// 	// Regular expression to match all "data:" occurrences
// 	re := regexp.MustCompile(`data:(\{.*?\})`)
// 	matches := re.FindAllStringSubmatch(string(body), -1)

// 	lastData := ""
// 	if len(matches) > 0 {
// 		lastData = matches[len(matches)-1][1]
// 	} else {
// 		fmt.Println("tidak ada data yang ditemukan")
// 	}

// 	if lastData != "" {
// 		lastData = strings.TrimSpace(lastData + "]}}]}}")
// 	}

// 	output := gjson.Get(lastData, "output.choices.0.message.content.0.text").String()
// 	_ = setToRedis(ctx.Request.Context(), req.Image, output, time.Second*10, handler.redisClient)

// 	ctx.JSON(http.StatusOK, helper.APIResponse(http.StatusOK, map[string]any{
// 		"output": output,
// 	}))
// }
