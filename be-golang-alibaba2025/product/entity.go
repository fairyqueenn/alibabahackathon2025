package product

import "encoding/json"

type (
	Product struct {
		ID                            int             `json:"id"`
		Name                          string          `json:"name"`
		Type                          string          `json:"type"`
		Price                         int             `json:"price"`
		Image                         string          `json:"image"`
		Description                   string          `json:"description"`
		Score                         int             `json:"score"`
		Grade                         string          `json:"grade"`
		Ingredients                   json.RawMessage `gorm:"type:json" json:"ingredients"`
		Nutritions                    json.RawMessage `gorm:"type:json" json:"nutritions"`
		IsForHalal                    bool            `json:"is_for_halal"`
		IsForVegan                    bool            `json:"is_for_vegan"`
		FriendlyForDiabetes           bool            `json:"friendly_for_diabetes"`
		FriendlyForHypertension       bool            `json:"friendly_for_hypertension"`
		FriendlyForCholesterol        bool            `json:"friendly_for_cholesterol"`
		FriendlyForGout               bool            `json:"friendly_for_gout"`
		FriendlyForKidney             bool            `json:"friendly_for_kidney"`
		FriendlyForCeliac             bool            `json:"friendly_for_celiac"`
		FriendlyForLactoseIntolerance bool            `json:"friendly_for_lactose_intolerance"`
		FriendlyForAutoimmune         bool            `json:"friendly_for_autoimmune"`
		FriendlyForThyroid            bool            `json:"friendly_for_thyroid"`
		FriendlyForObesity            bool            `json:"friendly_for_obesity"`
		FriendlyForDigestiveIssues    bool            `json:"friendly_for_digestive_issues"`
	}
)
