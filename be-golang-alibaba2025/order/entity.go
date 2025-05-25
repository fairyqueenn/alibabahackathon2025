package order

type Order struct {
	ID             int    `json:"id"`
	UserID         int    `json:"user_id"`
	ProductID      int    `json:"product_id"`
	Qty            int    `json:"qty"`
	Recommendation string `json:"recommendation"`
}
