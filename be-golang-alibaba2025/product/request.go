package product

import "encoding/json"

type (
	ProductUploadedFile struct {
		FileName string
		MIMEType string
		Size     int64
		Content  []byte
	}

	AddProductRequest struct {
		Name        string               `json:"name" form:"name"`
		Type        string               `json:"type" form:"type"`
		Price       int                  `json:"price" form:"price"`
		Image       *ProductUploadedFile `json:"-"`
		Description string               `json:"description" form:"description"`
	}

	UpdateProductRequest struct {
		Name        string          `json:"name"`
		Price       int             `json:"price"`
		Nutritions  json.RawMessage `json:"ingredients"`
		Description string          `json:"description"`
	}
)
