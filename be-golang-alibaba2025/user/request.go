package user

type (
	UpdateUserPreferenceRequest struct {
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
)
