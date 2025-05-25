package product

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/ZihxS/be-alibabacloud-genai-2025/constant"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/tidwall/gjson"
)

func (svc *service) GetAll(ctx context.Context) (products []*Product, err error) {
	return svc.repo.FindAll(ctx)
}

func (svc *service) GetByID(ctx context.Context, id int) (product *Product, err error) {
	return svc.repo.FindByID(ctx, id)
}

func (svc *service) AddProduct(ctx context.Context, req AddProductRequest) (product *Product, err error) {
	client, err := oss.New(constant.OSS_ENDPOINT, constant.OSS_ACCESS_KEY, constant.OSS_SECRET_KEY)
	if err != nil {
		return
	}

	bucket, err := client.Bucket(constant.OSS_BUCKET_NAME)
	if err != nil {
		return
	}

	expires := int64(60 * 5)
	signedURL, err := bucket.SignURL(req.Image.FileName, oss.HTTPGet, expires)
	if err != nil {
		return
	}

	urlImageReviewer := "http://8.215.3.103:8000/nutrition_scoring"

	requestBody := map[string]any{
		"image_url": signedURL,
		"query":     fmt.Sprintf("Buat output hanya json dan pakai bahasa indonesia untuk outputnya, output json harus json minify, hindari tambahan karakter dan new line juga hindari output \"```json\" dan \"```\". Berdasarkan gambar yang dicantumkan please follow JSON schema untuk arahan lebih lanjut. <JSONSchema>{\"description\":\"buat list ingredient dasar untuk membuat sebuah raw %s (deskripsi: %s) dan berikan takaran pasti untuk setiap ingredient nya serta menghitung nutrition bredasarkan ingredient serta menentukan type food atau beverage\",\"type\":\"object\",\"properties\":{\"type\":{\"description\":\"%s termasuk dalam food atau beverage hanya bisa memiliki value \"food\" dan \"beverage\" aja\",\"type\":\"string\",\"enum\":[\"food\",\"beverage\"]},\"ingredient\":{\"description\":\"buat list ingredient dasar untuk membuat sebuah raw %s (deskripsi: %s) dan berikan takaran pasti untuk setiap ingredient nya\",\"type\":\"array\",\"items\":{\"description\":\"List of Ingredients in English\",\"type\":\"object\",\"properties\":{\"name\":{\"description\":\"name of ingredients\",\"type\":\"string\"},\"measurements\":{\"description\":\"measurements of ingredients\",\"type\":\"string\"}},\"required\":[\"name\",\"measurements\"]}},\"nutrition\":{\"description\":\"jangan ada field tambahan\",\"type\":\"object\",\"properties\":{\"energy\":{\"description\":\"energy of ingredients in (KJ/100g for food) or (kcal/100ml for beverage)\",\"type\":\"number\"},\"saturated_fatty\":{\"description\":\"saturated fatty of ingredients in  (g/100g for food) or (g/100ml for beverage)\",\"type\":\"number\"},\"sugar\":{\"description\":\"sugar of ingredients in (g/100g for food) or (g/100ml for beverage)\",\"type\":\"number\"},\"salt\":{\"description\":\"salt of ingredients in (mg/100g)\",\"type\":\"number\"},\"protein\":{\"description\":\"protein of ingredients in (g/100g)\",\"type\":\"number\"},\"fibres\":{\"description\":\"fibers of ingredients in (g/100g)\",\"type\":\"number\"},\"fruit_vegetable_legumes\":{\"description\":\"persentase dari gabungan fruit, vegetables dan legumes dari ingredient dalam satuan (%%)\",\"type\":\"number\"}},\"required\":[\"energy\",\"saturated_fatty\",\"sugar\",\"salt\",\"protein\",\"fibres\",\"fruit_vegetable_legumes\"]}},\"required\":[\"type\",\"ingredient\",\"nutrition\"]}</JSONSchema>", req.Name, req.Description, req.Name, req.Name, req.Description),
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalf("error marshaling JSON: %v", err)
	}

	request, err := http.NewRequest("POST", urlImageReviewer, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("error creating request: %v", err)
	}

	request.Header.Add("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(request)
	if err != nil {
		log.Fatalf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading response body: %v", err)
	}

	text := gjson.Get(string(body), "response").String()
	ingredients := gjson.Parse(text)
	fmt.Printf("ingredients: %+v\n", text)

	parsedIngredients := ingredientsParser(ingredients)
	fmt.Printf("parsedIngredients: %+v\n", parsedIngredients)

	q := fmt.Sprintf("Buat output hanya json dan pakai bahasa indonesia untuk outputnya, output json harus json minify, hindari tambahan karakter dan new line juga hindari output \"```json\" dan \"```\". Berdasarkan gambar yang dicantumkan please follow JSON schema untuk arahan lebih lanjut. <JSONSchema>\"{\"description\":\"berikan perkiraan nutrition berdasarkan %s\",\"type\":\"object\",\"properties\":{\"nutrition\":{\"description\":\"dont add more field, only key and value and please answer with english\",\"type\":\"object\",\"properties\":{\"energy\":{\"description\":\"energy of ingredients in (KJ/100g for food) or (kcal/100ml for beverage)\",\"type\":\"number\"},\"saturated_fatty\":{\"description\":\"saturated fatty of ingredients in (g/100g for food) or (g/100ml for beverage)\",\"type\":\"number\"},\"sugar\":{\"description\":\"sugar of ingredients in (g/100g for food) or (g/100ml for beverage)\",\"type\":\"number\"},\"salt\":{\"description\":\"salt of ingredients in (mg/100g)\",\"type\":\"number\"},\"protein\":{\"description\":\"protein of ingredients in (g/100g)\",\"type\":\"number\"},\"fibres\":{\"description\":\"fibers of ingredients in (g/100g)\",\"type\":\"number\"},\"fruit_vegetable_legumes\":{\"description\":\"persentase dari gabungan fruit, vegetables dan legumes dari ingredient dalam satuan (%%)\",\"type\":\"number\"}},\"required\":[\"energy\",\"saturated_fatty\",\"sugar\",\"salt\",\"protein\",\"fibres\",\"fruit_vegetable_legumes\"]}},\"required\":[\"nutrition\"]}\"</JSONSchema>", parsedIngredients)
	requestBody = map[string]any{
		"image_url": signedURL,
		"query":     q,
	}

	fmt.Printf("q: %+v\n", q)

	jsonData, err = json.Marshal(requestBody)
	if err != nil {
		log.Fatalf("error marshaling JSON: %v", err)
	}

	request, err = http.NewRequest("POST", urlImageReviewer, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("error creating request: %v", err)
	}

	request.Header.Add("Content-Type", "application/json")

	httpClient = &http.Client{}
	resp, err = httpClient.Do(request)
	if err != nil {
		log.Fatalf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading response body: %v", err)
	}

	text = gjson.Get(string(body), "response").String()
	nutritions := gjson.Parse(text)
	_ = nutritions

	fmt.Printf("text: %+v\n", text)

	var n constant.Nutrition
	if err := json.Unmarshal([]byte(text), &n); err != nil {
		panic(err)
	}

	nType := nutritions.Get("type").String()
	fmt.Printf("nType: %+v\n", nType)

	if nType == "" {
		nType = ingredients.Get("type").String()
	}

	// call ai untuk menghasilkan ingredients
	// calculate nutri-score
	var c map[string]interface{}
	if nType == "food" {
		c, err = generateContent(nType, n, constant.FoodPoints)
		if err != nil {
			return
		}
	} else {
		c, err = generateContent(nType, n, constant.BeveragePoints)
		if err != nil {
			return
		}
	}
	// find grad
	// call ai untuk menghasilkan kalori
	// call ai untuk menghasilkan score packaging dan grade 1-5 for packaging sustainability dan info apakah bisa di daur ulang atau dipake lagi

	var s = float64(0)
	if x, ok := c["score"].(float64); ok {
		s = x
	}

	var g = ""
	if x, ok := c["grade"].(string); ok {
		g = x
	}

	product = &Product{
		Name:        req.Name,
		Type:        nType,
		Price:       req.Price,
		Image:       req.Image.FileName,
		Description: req.Description,
		Score:       int(s),                           // dari AI
		Grade:       g,                                // dari AI
		Ingredients: json.RawMessage(ingredients.Raw), // dari AI
		Nutritions:  json.RawMessage(nutritions.Raw),  // dari AI
		IsForHalal:  true,
	}

	err = svc.repo.UpsertProduct(ctx, product)
	if err != nil {
		return
	}

	// go routine: call ai untuk kategori kesehatan

	return
}

func (svc *service) UpdateProduct(ctx context.Context, id int, req UpdateProductRequest) (err error) {
	product, err := svc.repo.FindByID(ctx, id)
	if err != nil {
		return
	}

	product.Name = req.Name
	product.Price = req.Price
	product.Description = req.Description
	product.Nutritions = req.Nutritions

	return svc.repo.UpsertProduct(ctx, product)
}

func ingredientsParser(ingredients gjson.Result) string {
	var parts []string

	ingredients.Get("ingredient").ForEach(func(key, value gjson.Result) bool {
		name := value.Get("name").String()
		measurements := value.Get("measurements").String()
		if name != "" && measurements != "" {
			parts = append(parts, name+" "+measurements)
		}
		return true
	})

	return strings.Join(parts, ", ")
}

type Payload struct {
	Energy                int
	SaturatedFatty        int
	Sugar                 int
	Salt                  int
	Protein               int
	Fibres                int
	FruitVegetableLegumes int
}

func calculatePointFood(payload constant.Nutrition) float64 {
	N := payload.Energy + payload.SaturatedFatty + payload.Sugar + payload.Salt
	P := float64(0)

	if N >= 11 {
		P += payload.Fibres + payload.FruitVegetableLegumes
	} else {
		P += payload.Protein + payload.Fibres + payload.FruitVegetableLegumes
	}

	return N - P
}

func calculatePointBeverage(payload constant.Nutrition) float64 {
	N := payload.Energy + payload.SaturatedFatty + payload.Sugar + payload.Salt
	P := payload.Protein + payload.Fibres + payload.FruitVegetableLegumes

	return N - P
}

func generateContent(itemType string, nutrition constant.Nutrition, points constant.NutrientPoints) (map[string]interface{}, error) {
	var checkGrade func(int) string

	calculate := "food"
	switch itemType {
	case "food":
		calculate = "food"
		checkGrade = checkGradeFood
	case "beverage":
		calculate = "age"
		checkGrade = checkGradeBeverage
	default:
		return nil, fmt.Errorf("generated failed, not valid")
	}

	allPoint := map[string]int{
		"energy":                  constant.GetPoint(nutrition.Energy, points.Energy),
		"sugar":                   constant.GetPoint(nutrition.Sugar, points.Sugar),
		"saturated_fatty":         constant.GetPoint(nutrition.SaturatedFatty, points.SaturatedFatty),
		"salt":                    constant.GetPoint(nutrition.Salt, points.Salt),
		"protein":                 constant.GetPoint(nutrition.Protein, points.Protein),
		"fibres":                  constant.GetPoint(nutrition.Fibres, points.Fibres),
		"fruit_vegetable_legumes": constant.GetPoint(nutrition.FruitVegetableLegumes, points.FruitVegetableLegumes),
	}

	nutritionForCalc := constant.Nutrition{
		Energy:                float64(allPoint["energy"]),
		Sugar:                 float64(allPoint["sugar"]),
		SaturatedFatty:        float64(allPoint["saturated_fatty"]),
		Salt:                  float64(allPoint["salt"]),
		Protein:               float64(allPoint["protein"]),
		Fibres:                float64(allPoint["fibres"]),
		FruitVegetableLegumes: float64(allPoint["fruit_vegetable_legumes"]),
	}

	score := calculatePointFood(nutritionForCalc)
	if calculate == "bevarage" {
		score = calculatePointBeverage(nutritionForCalc)
	}
	grade := checkGrade(int(score))

	return map[string]interface{}{
		"score": score,
		"grade": grade,
	}, nil
}

func checkGradeFood(point int) string {
	switch {
	case point <= 0:
		return "A"
	case point <= 2:
		return "B"
	case point <= 10:
		return "E"
	case point <= 18:
		return "D"
	default:
		return "E"
	}
}

func checkGradeBeverage(point int) string {
	switch {
	case point <= 0:
		return "A"
	case point <= 2:
		return "B"
	case point <= 6:
		return "C"
	case point <= 9:
		return "D"
	default:
		return "E"
	}
}
