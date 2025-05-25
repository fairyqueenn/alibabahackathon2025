package user

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserPreference struct {
	ID                   int  `json:"id"`
	UserID               int  `json:"user_id"`
	IsHalal              bool `json:"is_halal"`
	IsVegan              bool `json:"is_vegan"`
	IsDiabetes           bool `json:"is_diabetes"`
	IsHypertension       bool `json:"is_hypertension"`
	IsCholesterol        bool `json:"is_cholesterol"`
	IsGout               bool `json:"is_gout"`
	IsKidney             bool `json:"is_kidney"`
	IsCeliac             bool `json:"is_celiac"`
	IsLactoseIntolerance bool `json:"is_lactose_intolerance"`
	IsAutoimmune         bool `json:"is_autoimmune"`
	IsThyroid            bool `json:"is_thyroid"`
	IsObesity            bool `json:"is_obesity"`
	IsDigestiveIssues    bool `json:"is_digestive_issues"`
}

type UserRecommendation struct {
	ID             int    `json:"id"`
	UserID         int    `json:"user_id"`
	Recommendation string `json:"recommendation"`
}
